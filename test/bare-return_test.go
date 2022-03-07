package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

func TestBareReturn(t *testing.T) {
	testRule(t, "bare-return", &rule.BareReturnRule{})
}
