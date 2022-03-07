package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

func TestGetReturn(t *testing.T) {
	testRule(t, "get-return", &rule.GetReturnRule{})
}
