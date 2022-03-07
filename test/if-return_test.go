package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// TestIfReturn rule.
func TestIfReturn(t *testing.T) {
	testRule(t, "if-return", &rule.IfReturnRule{})
}
