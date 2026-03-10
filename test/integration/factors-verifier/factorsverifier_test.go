package isevenverifier_test

import (
	_ "embed"
	"encoding/binary"
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/tahardi/bearchain/contracts/bindings"
	"github.com/tahardi/bearchain/test/foundry"
	"github.com/tahardi/bearchain/test/integration"
)

const (
	ContractName = "FactorsVerifier"
)

//go:embed testdata/groth16.json
var groth16Bytes []byte

type Groth16Proof struct {
	ImageID []byte `json:"image_id"`
	Journal []byte `json:"journal"`
	Seal    []byte `json:"seal"`
}

func groth16ProofFromTestData(t *testing.T) Groth16Proof {
	t.Helper()
	groth16Proof := Groth16Proof{}
	err := json.Unmarshal(groth16Bytes, &groth16Proof)
	require.NoError(t, err)
	return groth16Proof
}

func deployContract(
	t *testing.T,
	anvil *foundry.Anvil,
	owner *foundry.Account,
) *bindings.FactorsVerifier {
	t.Helper()
	contractAddress, err := anvil.DeployContract(t.Context(), ContractName, owner)
	require.NoError(t, err)

	client, err := anvil.Client()
	require.NoError(t, err)

	contract, err := bindings.NewFactorsVerifier(*contractAddress, client)
	require.NoError(t, err)
	return contract
}

func newTransactionOpts(
	t *testing.T,
	anvil *foundry.Anvil,
	from *foundry.Account,
) *bind.TransactOpts {
	t.Helper()
	opts, err := bind.NewKeyedTransactorWithChainID(from.PrivateKey(), anvil.ChainID())
	require.NoError(t, err)
	return opts
}

func executeCall(
	t *testing.T,
	anvil *foundry.Anvil,
	contractCall func() (*types.Transaction, error),
) (*types.Receipt, error) {
	t.Helper()
	tx, err := contractCall()
	if err != nil {
		return nil, err
	}

	client, err := anvil.Client()
	if err != nil {
		return nil, err
	}
	return bind.WaitMined(t.Context(), client, tx)
}

func verify(
	t *testing.T,
	anvil *foundry.Anvil,
	contract *bindings.FactorsVerifier,
	account *foundry.Account,
	product uint64,
	seal []byte,
) (*types.Receipt, error) {
	t.Helper()
	opts := newTransactionOpts(t, anvil, account)
	call := func() (*types.Transaction, error) {
		return contract.Verify(opts, product, seal)
	}
	return executeCall(t, anvil, call)
}

func TestFactorsVerifier_Verify(t *testing.T) {
	// given
	anvil, stop := integration.StartAnvil(t, true)
	defer stop()

	owner := anvil.Account(0)
	contract := deployContract(t, anvil, owner)

	// Our Rust ZKVM program outputs the journal (i.e., our expected product)
	// in Big Endian format. If you want the actual value, convert to LE.
	groth16Proof := groth16ProofFromTestData(t)
	product := binary.BigEndian.Uint64(groth16Proof.Journal)

	// when
	_, err := verify(t, anvil, contract, owner, product, groth16Proof.Seal)

	// then
	require.NoError(t, err)
}
