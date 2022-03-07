package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// TestTimeNamingRule rule.
func TestTimeNaming(t *testing.T) {
	testRule(t, "time-naming", &rule.TimeNamingRule{})
}
