package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// BoolLiteral rule.
func TestBoolLiteral(t *testing.T) {
	testRule(t, "bool-literal-in-expr", &rule.BoolLiteralRule{})
}
