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
        bytes32 zkvmImageId = hex"2bda51ae4f0326636e89480a383a770f713d514f76f6d58a5c2d1373b6b87d48";
        //        bytes32 zkvmImageId = vm.envBytes32("ZK_VM_IMAGE_ID");

        vm.startBroadcast();
        IRiscZeroVerifier riscZeroVerifier =
            new RiscZeroGroth16Verifier(ControlID.CONTROL_ROOT, ControlID.BN254_CONTROL_ID);
        factorsVerifier = new FactorsVerifier(riscZeroVerifier, zkvmImageId);
        vm.stopBroadcast();
    }
}
