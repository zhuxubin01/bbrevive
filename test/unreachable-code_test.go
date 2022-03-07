package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

func TestUnreachableCode(t *testing.T) {
	testRule(t, "unreachable-code", &rule.UnreachableCodeRule{})
}
