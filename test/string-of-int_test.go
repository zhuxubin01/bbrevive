package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// String-of-int rule.
func TestStringOfInt(t *testing.T) {
	testRule(t, "string-of-int", &rule.StringOfIntRule{})
}
