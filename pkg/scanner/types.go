package scanner

type Severity string

const (
	SeverityCritical Severity = "CRITICAL"
	SeverityHigh     Severity = "HIGH"
	SeverityMedium   Severity = "MEDIUM"
	SeverityLow      Severity = "LOW"
	SeverityInfo     Severity = "INFO"
)

type Issue struct {
	RuleID      string   `json:"rule_id"`
	RuleName    string   `json:"rule_name"`
	Severity    Severity `json:"severity"`
	File        string   `json:"file"`
	Line        int      `json:"line"`
	Snippet     string   `json:"snippet"`
	Remediation string   `json:"remediation"`
}

type ScanResult struct {
	TotalFilesScanned int     `json:"total_files_scanned"`
	DurationMs        int64   `json:"duration_ms"`
	Issues            []Issue `json:"issues"`
}
