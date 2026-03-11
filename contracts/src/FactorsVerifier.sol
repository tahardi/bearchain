pragma solidity ^0.8.33;

import {IRiscZeroVerifier} from "@risc0-ethereum/IRiscZeroVerifier.sol";

// This contract verifies RISC Zero proofs that prove the knowledge of some
// factors a and b that when multiplied together produce c.
contract FactorsVerifier {
    address public immutable OWNER;
    bytes32 public imageId;
    IRiscZeroVerifier public immutable RISC_ZERO_VERIFIER;

    event FactorsKnownForProduct(uint256 indexed product);
    event ImageIdUpdated(bytes32 indexed newImageId);
    event DebugVerification(bytes4 sealSelector, bytes4 expectedSelector);

    constructor(IRiscZeroVerifier _riscZeroVerifier, bytes32 _imageId) {
        OWNER = msg.sender;
        imageId = _imageId;
        RISC_ZERO_VERIFIER = _riscZeroVerifier;
    }

    function setImageId(bytes32 _imageId) external {
        require(msg.sender == OWNER, "Only owner can update imageID");
        imageId = _imageId;
        emit ImageIdUpdated(_imageId);
    }

    function verify(uint64 product, bytes calldata seal) external {
        // Must use packed format or else the product will not match the one
        // in the seal. Consider updating ZKVM program to output non-packed val.
        bytes memory journal = abi.encodePacked(product);
        RISC_ZERO_VERIFIER.verify(seal, imageId, sha256(journal));
        emit FactorsKnownForProduct(product);
    }
}
