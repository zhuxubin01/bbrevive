package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// TestModifiesParam rule.
func TestModifiesParam(t *testing.T) {
	testRule(t, "modifies-param", &rule.ModifiesParamRule{})
}
