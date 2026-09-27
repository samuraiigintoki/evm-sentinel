// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract VulnerableVault {
    address public owner;

    constructor() {
        owner = msg.sender;
    }

    function transferOwnership(address newOwner) public {
        require(tx.origin == owner, "Only owner");
        owner = newOwner;
    }

    function emergencyDrain(address payable recipient) public {
        require(msg.sender == owner, "Only owner");
        selfdestruct(recipient);
    }
}