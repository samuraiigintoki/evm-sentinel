package scanner

import "regexp"

// AdvancedRules adds detection for complex DeFi attack vectors.
func AdvancedRules() []Rule {
	return []Rule{
		{
			ID:          "EVM-006",
			Name:        "Potential Read-Only Reentrancy Query",
			Severity:    SeverityMedium,
			Regex:       regexp.MustCompile(`getPoolState|getVirtualPrice|calculatePrice`),
			Remediation: "Ensure external view/query functions guard against intermediate transient state manipulation.",
		},
		{
			ID:          "EVM-007",
			Name:        "Unprotected Token Transfer From Arbitrary Address",
			Severity:    SeverityHigh,
			Regex:       regexp.MustCompile(`transferFrom\s*\(\s*[^,]+,\s*address\(this\)`),
			Remediation: "Verify spender allowance and validate the source address before pulling funds.",
		},
	}
}
