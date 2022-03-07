package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

func TestDuplicatedImports(t *testing.T) {
	testRule(t, "duplicated-imports", &rule.DuplicatedImportsRule{})
}
