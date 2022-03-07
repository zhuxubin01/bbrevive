package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

// TestStructTag tests struct-tag rule
func TestStructTag(t *testing.T) {
	testRule(t, "struct-tag", &rule.StructTagRule{})
}
