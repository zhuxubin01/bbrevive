package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

func TestWaitGroupByValue(t *testing.T) {
	testRule(t, "waitgroup-by-value", &rule.WaitGroupByValueRule{})
}
