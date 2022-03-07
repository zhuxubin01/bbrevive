package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

func TestUnexportednaming(t *testing.T) {
	testRule(t, "unexported-naming", &rule.UnexportedNamingRule{})
}
