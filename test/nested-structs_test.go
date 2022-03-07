package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/lint"
	"github.com/zhuxubin01/bbrevive/rule"
)

func TestNestedStructs(t *testing.T) {
	testRule(t, "nested-structs", &rule.NestedStructs{}, &lint.RuleConfig{})
}
