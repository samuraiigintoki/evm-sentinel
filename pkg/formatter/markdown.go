package formatter

import (
	"fmt"
	"strings"

	"evm-sentinel/pkg/scanner"
)

// GenerateMarkdownReport converts scan results into a GitHub-compatible markdown table.
func GenerateMarkdownReport(res *scanner.ScanResult) string {
	var sb strings.Builder

	sb.WriteString("# 🛡️ EVM-Sentinel Security Audit Summary\n\n")
	sb.WriteString(fmt.Sprintf("- **Total Files Scanned:** `%d`\n", res.TotalFilesScanned))
	sb.WriteString(fmt.Sprintf("- **Scan Duration:** `%dms`\n", res.DurationMs))
	sb.WriteString(fmt.Sprintf("- **Vulnerabilities Identified:** `%d`\n\n", len(res.Issues)))

	if len(res.Issues) == 0 {
		sb.WriteString("✅ **Clean Audit:** No static vulnerabilities found in target smart contracts.\n")
		return sb.String()
	}

	sb.WriteString("| # | Severity | Rule ID | Description | File Location |\n")
	sb.WriteString("|:---|:---|:---|:---|:---|\n")

	for i, issue := range res.Issues {
		sb.WriteString(fmt.Sprintf("| %d | **%s** | `%s` | %s | `%s:%d` |\n",
			i+1, issue.Severity, issue.RuleID, issue.RuleName, issue.File, issue.Line))
	}

	return sb.String()
}
