package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// TestSuperfluousElse rule.
func TestSuperfluousElse(t *testing.T) {
	testRule(t, "superfluous-else", &rule.SuperfluousElseRule{})
}
