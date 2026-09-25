package scanner

import "regexp"

func DefaultRules() []Rule {
	return []Rule{
		{
			ID:          "EVM-001",
			Name:        "tx.origin Authentication Vector",
			Severity:    SeverityHigh,
			Regex:       regexp.MustCompile(`require\s*\(\s*tx\.origin\s*==`),
			Remediation: "Use msg.sender for authorization to avoid phishing call vulnerabilities.",
		},
		{
			ID:          "EVM-002",
			Name:        "Unchecked Low-Level Call",
			Severity:    SeverityHigh,
			Regex:       regexp.MustCompile(`\.call\{value:`),
			Remediation: "Check the boolean return value of low-level calls explicitly.",
		},
		{
			ID:          "EVM-003",
			Name:        "Arbitrary Delegatecall Target",
			Severity:    SeverityCritical,
			Regex:       regexp.MustCompile(`\.(delegatecall)\(`),
			Remediation: "Ensure target address is immutable or checked against an authorized whitelist.",
		},
		{
			ID:          "EVM-004",
			Name:        "Dangerous Selfdestruct",
			Severity:    SeverityCritical,
			Regex:       regexp.MustCompile(`selfdestruct\s*\(`),
			Remediation: "Avoid selfdestruct opcode; consider contract pause mechanics instead.",
		},
		{
			ID:          "EVM-005",
			Name:        "Block Timestamp Manipulation",
			Severity:    SeverityLow,
			Regex:       regexp.MustCompile(`block\.timestamp`),
			Remediation: "Do not use block.timestamp as a source of randomness or strict time gates.",
		},
	}
}