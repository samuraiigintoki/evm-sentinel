package scanner

import (
	"testing"
)

func TestEngineRuleMatching(t *testing.T) {
	rules := DefaultRules()
	ruleMap := make(map[string]Rule)
	for _, r := range rules {
		ruleMap[r.ID] = r
	}

	tests := []struct {
		name      string
		ruleID    string
		inputLine string
		shouldHit bool
	}{
		{"tx.origin check", "EVM-001", "require(tx.origin == owner, 'Unauthorized');", true},
		{"safe sender check", "EVM-001", "require(msg.sender == owner, 'Unauthorized');", false},
		{"delegatecall check", "EVM-003", "target.delegatecall(data);", true},
		{"selfdestruct check", "EVM-004", "selfdestruct(payable(recipient));", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, exists := ruleMap[tt.ruleID]
			if !exists {
				t.Fatalf("rule %s not found", tt.ruleID)
			}
			matched := r.Regex.MatchString(tt.inputLine)
			if matched != tt.shouldHit {
				t.Errorf("expected match=%v, got=%v for snippet: %s", tt.shouldHit, matched, tt.inputLine)
			}
		})
	}
}
