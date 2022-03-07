package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// TestEarlyReturn tests early-return rule.
func TestEarlyReturn(t *testing.T) {
	testRule(t, "early-return", &rule.EarlyReturnRule{})
}
