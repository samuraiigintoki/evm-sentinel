# Contributing to EVM-Sentinel

We welcome contributions from Web3 security researchers and Go developers!

## Adding New Security Rules
1. Open `pkg/scanner/rules.go` or `pkg/scanner/rules_advanced.go`.
2. Define a new `Rule` struct with:
   - Unique ID (e.g., `EVM-008`)
   - Severity level (`Critical`, `High`, `Medium`, `Low`)
   - Optimized regex pattern
   - Actionable remediation advice
3. Add corresponding test cases in `pkg/scanner/rules_test.go`.
4. Run tests: `make test`

## Code Standards
- Adhere to idiomatic Go naming and formatting (`go fmt ./...`).
- Zero external third-party dependencies required for core scanner engine.
