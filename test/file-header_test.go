package test

import (
	"testing"

	"github.com/zhuxubin01/bbrevive/lint"
	"github.com/zhuxubin01/bbrevive/rule"
)

func TestLintFileHeader(t *testing.T) {
	testRule(t, "lint-file-header1", &rule.FileHeaderRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"foobar"},
	})

	testRule(t, "lint-file-header2", &rule.FileHeaderRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"foobar"},
	})

	testRule(t, "lint-file-header3", &rule.FileHeaderRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"foobar"},
	})

	testRule(t, "lint-file-header4", &rule.FileHeaderRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"^\\sfoobar$"},
	})

	testRule(t, "lint-file-header5", &rule.FileHeaderRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"^\\sfoo.*bar$"},
	})

	testRule(t, "lint-file-header6", &rule.FileHeaderRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"foobar"},
	})
}

func BenchmarkLintFileHeader(b *testing.B) {
	var t *testing.T
	for i := 0; i <= b.N; i++ {
		testRule(t, "lint-file-header1", &rule.FileHeaderRule{}, &lint.RuleConfig{
			Arguments: []interface{}{"foobar"},
		})
	}
}
