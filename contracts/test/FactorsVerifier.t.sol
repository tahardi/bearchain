pragma solidity ^0.8.33;

import {Test} from "@forge-std/Test.sol";
import {RiscZeroCheats} from "@risc0-ethereum/test/RiscZeroCheats.sol";
import {Receipt as RiscZeroReceipt} from "@risc0-ethereum/IRiscZeroVerifier.sol";
import {VerificationFailed} from "@risc0-ethereum/IRiscZeroVerifier.sol";
import {RiscZeroMockVerifier} from "@risc0-ethereum/test/RiscZeroMockVerifier.sol";
import {FactorsVerifier} from "../src/FactorsVerifier.sol";

contract FactorsVerifierTest is RiscZeroCheats, Test {
    address public owner;
    bytes32 public constant IMAGE_ID = 0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef;

    FactorsVerifier public factorsVerifier;
    RiscZeroMockVerifier public riscZeroVerifier;

    function setUp() public {
        owner = address(1);
        riscZeroVerifier = new RiscZeroMockVerifier(0);

        vm.prank(owner);
        factorsVerifier = new FactorsVerifier(riscZeroVerifier, IMAGE_ID);
    }

    function test_verify() public {
        // given
        uint64 product = 4;
        RiscZeroReceipt memory receipt = riscZeroVerifier.mockProve(IMAGE_ID, sha256(abi.encodePacked(product)));

        vm.prank(owner);
        vm.expectEmit(true, false, false, true);
        emit FactorsVerifier.FactorsKnownForProduct(product);

        // when/then
        factorsVerifier.verify(product, receipt.seal);
    }

    function test_verify_wrong_product() public {
        // given
        uint64 product = 4;
        uint64 wrongProduct = 5;
        RiscZeroReceipt memory receipt = riscZeroVerifier.mockProve(IMAGE_ID, sha256(abi.encodePacked(product)));

        vm.prank(owner);
        vm.expectRevert(VerificationFailed.selector);

        // when/then
        factorsVerifier.verify(wrongProduct, receipt.seal);
    }
}
