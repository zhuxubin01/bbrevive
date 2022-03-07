package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

func TestUnusedReceiver(t *testing.T) {
	testRule(t, "unused-receiver", &rule.UnusedReceiverRule{})
}
