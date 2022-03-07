package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// TestUnnecessaryStmt rule.
func TestUnnecessaryStmt(t *testing.T) {
	testRule(t, "unnecessary-stmt", &rule.UnnecessaryStmtRule{})
}
