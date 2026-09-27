package scanner

import "regexp"

func DefaultRules() []Rule {
	return []Rule{
		{
			ID:          "EVM-001",
			Name:        "tx.origin Authorization Vector",
			Severity:    SeverityHigh,
			Regex:       regexp.MustCompile(`require\s*\(\s*tx\.origin\s*==`),
			Remediation: "Use msg.sender instead of tx.origin to prevent phishing and unauthorized access.",
		},
		{
			ID:          "EVM-002",
			Name:        "Reentrancy / Unchecked Call Vector",
			Severity:    SeverityCritical,
			Regex:       regexp.MustCompile(`\.call\{value:`),
			Remediation: "Follow the Checks-Effects-Interactions (CEI) pattern or use OpenZeppelin ReentrancyGuard.",
		},
		{
			ID:          "EVM-003",
			Name:        "Arbitrary Delegatecall Proxy Injection",
			Severity:    SeverityCritical,
			Regex:       regexp.MustCompile(`\.(delegatecall)\(`),
			Remediation: "Ensure target address for delegatecall is strictly immutable or whitelisted.",
		},
		{
			ID:          "EVM-004",
			Name:        "AMM Spot Price Oracle Manipulation",
			Severity:    SeverityHigh,
			Regex:       regexp.MustCompile(`getReserves\s*\(\)`),
			Remediation: "Avoid spot prices from AMM pools vulnerable to flash loans; use TWAP or Chainlink feeds.",
		},
		{
			ID:          "EVM-005",
			Name:        "DoS via Unbounded Array Iteration",
			Severity:    SeverityMedium,
			Regex:       regexp.MustCompile(`for\s*\(.*;\s*.*<\s*\w+\.length;`),
			Remediation: "Avoid looping over dynamic storage arrays. Implement pull-over-push payment patterns.",
		},
	}
}