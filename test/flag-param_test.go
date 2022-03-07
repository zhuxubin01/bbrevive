package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

func TestFlagParam(t *testing.T) {
	testRule(t, "flag-param", &rule.FlagParamRule{})
}
