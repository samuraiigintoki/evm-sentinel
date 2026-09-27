// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

/**
 * @title VulnerableVault (Test Harness)
 * @dev Intentionally vulnerable mock contract for EVM-Sentinel static analysis tests.
 */
contract VulnerableVault {
    address public owner;

    constructor() {
        owner = msg.sender;
    }

    function transferOwnership(address newOwner) public {
        // Triggers EVM-001 (tx.origin authentication vector)
        require(tx.origin == owner, "Only owner");
        owner = newOwner;
    }

    function emergencyDrain(address payable recipient) public {
        require(msg.sender == owner, "Only owner");
        // Triggers EVM-004 (Dangerous selfdestruct)
        assembly {
            selfdestruct(recipient)
        }
    }
}