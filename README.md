# 🛡️ EVM-Sentinel

A high-performance static security analysis tool for Solidity smart contracts, built in Go.

## Features
- ⚡ **Concurrent Analysis:** Multi-threaded worker pool to scan multi-contract repos in milliseconds.
- 🎯 **Security Rules:** Built-in pattern detection for critical Web3 vulnerability classes.
- 📊 **Rich Output:** Clean terminal tabular reports and JSON export for CI/CD pipelines.

## Rule Matrix
| ID | Vulnerability Class | Severity | Description |
|:---|:--------------------|:---------|:------------|
| `EVM-001` | Phishing via `tx.origin` | High | Checks authorization relying on `tx.origin` instead of `msg.sender` |
| `EVM-002` | Unchecked Low-Level Call | High | Detects low-level `.call()` without checking return boolean |
| `EVM-003` | Dangerous `delegatecall` | Critical | Identifies untrusted proxy target invocations |
| `EVM-004` | Selfdestruct Usage | Critical | Flags use of deprecated and risky `selfdestruct` opcode |
| `EVM-005` | Timestamp Dependence | Low | Flags reliance on `block.timestamp` for game logic |

## Installation & Usage
```bash
go build -o evm-sentinel cmd/evmsentinel/main.go
./evm-sentinel --dir ./contracts