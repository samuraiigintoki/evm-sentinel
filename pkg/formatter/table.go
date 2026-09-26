package formatter

import (
	"fmt"

	"evm-sentinel/pkg/scanner"
)

func RenderConsoleReport(res *scanner.ScanResult) {
	fmt.Println("\n========================================================")
	fmt.Printf("🛡️  EVM-SENTINEL AUDIT REPORT | %d files scanned in %dms\n", res.TotalFilesScanned, res.DurationMs)
	fmt.Println("========================================================")

	if len(res.Issues) == 0 {
		fmt.Println("✅ No known static security vulnerabilities detected.")
		return
	}

	for i, issue := range res.Issues {
		fmt.Printf("[%d] [%s] %s\n", i+1, issue.Severity, issue.RuleName)
		fmt.Printf("    File:   %s:%d\n", issue.File, issue.Line)
		fmt.Printf("    Code:   %s\n", issue.Snippet)
		fmt.Printf("    Fix:    %s\n", issue.Remediation)
		fmt.Println("--------------------------------------------------------")
	}

	fmt.Printf("\nFound %d potential vulnerabilities.\n\n", len(res.Issues))
}