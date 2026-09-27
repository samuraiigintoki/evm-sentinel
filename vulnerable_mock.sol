// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IUniswapV2Pair {
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
}

/**
 * @title DeFiLendingVault (Test Harness)
 * @notice Demonstrates top Web3 vulnerability classes for EVM-Sentinel static analysis.
 */
contract DeFiLendingVault {
    address public owner;
    address public pricePair;
    address[] public stakeholders;
    mapping(address => uint256) public balances;

    constructor(address _pair) {
        owner = msg.sender;
        pricePair = _pair;
    }

    // 1. EVM-001: Phishing authorization vector
    function setPairAddress(address newPair) external {
        require(tx.origin == owner, "Unauthorized");
        pricePair = newPair;
    }

    // 2. EVM-002: Reentrancy vulnerability (State updated after call)
    function withdraw(uint256 amount) external {
        require(balances[msg.sender] >= amount, "Insufficient balance");
        
        (bool success, ) = msg.sender.call{value: amount}("");
        require(success, "Transfer failed");

        balances[msg.sender] -= amount;
    }

    // 3. EVM-003: Arbitrary delegatecall execution
    function executeDelegate(address target, bytes calldata data) external {
        (bool ok, ) = target.delegatecall(data);
        require(ok, "Delegatecall failed");
    }

    // 4. EVM-004: AMM Spot Price Oracle manipulation via Flash Loan
    function getCollateralPrice() public view returns (uint256) {
        (uint112 reserve0, uint112 reserve1, ) = IUniswapV2Pair(pricePair).getReserves();
        return (uint256(reserve1) * 1e18) / uint256(reserve0);
    }

    // 5. EVM-005: Denial of Service via unbounded loop over dynamic array
    function distributeYield() external {
        for (uint256 i = 0; i < stakeholders.length; i++) {
            (bool success, ) = payable(stakeholders[i]).call{value: 1 ether}("");
            require(success, "Transfer failed");   
        }
    }
}