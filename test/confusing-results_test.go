package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

func TestConfusingResults(t *testing.T) {
	testRule(t, "confusing-results", &rule.ConfusingResultsRule{})
}
