package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

func TestUnconditionalRecursion(t *testing.T) {
	testRule(t, "unconditional-recursion", &rule.UnconditionalRecursionRule{})
}
