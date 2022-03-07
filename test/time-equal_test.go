package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// TestTimeEqual rule.
func TestTimeEqual(t *testing.T) {
	testRule(t, "time-equal", &rule.TimeEqualRule{})
}
