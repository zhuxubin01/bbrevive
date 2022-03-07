package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/lint"
	"github.com/zhuxubin01/bbrevive/rule"
)

func TestAddConstant(t *testing.T) {
	args := []interface{}{map[string]interface{}{
		"maxLitCount": "2",
		"allowStrs":   "\"\"",
		"allowInts":   "0,1,2",
		"allowFloats": "0.0,1.0",
	}}

	testRule(t, "add-constant", &rule.AddConstantRule{}, &lint.RuleConfig{
		Arguments: args,
	})
}
