package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// UselessBreak rule.
func TestUselessBreak(t *testing.T) {
	testRule(t, "useless-break", &rule.UselessBreak{})
}
