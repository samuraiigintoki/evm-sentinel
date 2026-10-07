
package scanner

// Config defines execution parameters and rule filtering options for EVM-Sentinel.
type Config struct {
	ExcludedRules []string `json:"excluded_rules"`
	MinSeverity   Severity `json:"min_severity"`
	MaxWorkers    int      `json:"max_workers"`
	Directory     string   `json:"directory"`
}

// DefaultConfig returns safe baseline configuration values.
func DefaultConfig() *Config {
	return &Config{
		ExcludedRules: make([]string, 0),
		MinSeverity:   SeverityLow,
		MaxWorkers:    4,
		Directory:     ".",
	}
}
