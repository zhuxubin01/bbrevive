package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// TestConfusingNaming rule.
func TestConfusingNaming(t *testing.T) {
	testRule(t, "confusing-naming1", &rule.ConfusingNamingRule{})
}
