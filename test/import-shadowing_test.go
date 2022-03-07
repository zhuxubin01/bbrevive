package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/rule"
)

func TestImportShadowing(t *testing.T) {
	testRule(t, "import-shadowing", &rule.ImportShadowingRule{})
}
