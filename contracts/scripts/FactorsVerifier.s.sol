pragma solidity ^0.8.33;

import {FactorsVerifier} from "../src/FactorsVerifier.sol";
import {ControlID} from "@risc0-ethereum/groth16/ControlID.sol";
import {IRiscZeroVerifier} from "@risc0-ethereum/IRiscZeroVerifier.sol";
import {RiscZeroGroth16Verifier} from "@risc0-ethereum/groth16/RiscZeroGroth16Verifier.sol";
import {Script} from "@forge-std/Script.sol";

contract FactorsVerifierScript is Script {
    FactorsVerifier public factorsVerifier;

    function setUp() public {}

    function run() public {
        // This is taken from:
        // bearclave-contracts/test/integration/factors-verifier/testdata/groth16.json
        //
        // It represents a specific version of the bearclave-zkvm/factors-verifier
        // program. This must be updated if the groth16 testdata gets updated.
        // That, or the test should explicitly call setImageId to ensure the
        // right imageId is being used to verify a given Groth16 seal.
        bytes32 imageId = hex"2bda51ae4f0326636e89480a383a770f713d514f76f6d58a5c2d1373b6b87d48";

        vm.startBroadcast();
        IRiscZeroVerifier riscZeroVerifier =
            new RiscZeroGroth16Verifier(ControlID.CONTROL_ROOT, ControlID.BN254_CONTROL_ID);
        factorsVerifier = new FactorsVerifier(riscZeroVerifier, imageId);
        vm.stopBroadcast();
    }
}
