package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

func TestModifiesValRec(t *testing.T) {
	testRule(t, "modifies-value-receiver", &rule.ModifiesValRecRule{})
}
