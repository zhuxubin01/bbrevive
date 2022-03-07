package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// Atomic rule.
func TestAtomic(t *testing.T) {
	testRule(t, "atomic", &rule.AtomicRule{})
}
