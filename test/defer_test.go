package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/lint"
	"github.com/zhuxubin01/bbrevive/rule"
)

// Defer rule.
func TestDefer(t *testing.T) {
	testRule(t, "defer", &rule.DeferRule{})
}

func TestDeferLoopDisabled(t *testing.T) {
	testRule(t, "defer-loop-disabled", &rule.DeferRule{}, &lint.RuleConfig{
		Arguments: []interface{}{[]interface{}{"return", "recover", "call-chain", "method-call"}},
	})
}

func TestDeferOthersDisabled(t *testing.T) {
	testRule(t, "defer-only-loop-enabled", &rule.DeferRule{}, &lint.RuleConfig{
		Arguments: []interface{}{[]interface{}{"loop"}},
	})
}
