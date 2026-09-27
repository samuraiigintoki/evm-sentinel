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
		{"reentrancy call check", "EVM-002", "(bool s, ) = msg.sender.call{value: amount}(\"\");", true},
		{"delegatecall check", "EVM-003", "target.delegatecall(data);", true},
		{"spot price oracle check", "EVM-004", "(uint112 r0, uint112 r1, ) = pair.getReserves();", true},
		{"unbounded loop DoS check", "EVM-005", "for (uint256 i = 0; i < stakeholders.length; i++) {", true},
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