package test

import (
	"testing"

	"github.com/mgechev/revive/lint"
	"github.com/mgechev/revive/rule"
)

func TestFunctionSignatureStyle(t *testing.T) {
	testRule(t, "function-signature-style", &rule.FunctionSignatureStyleRule{}, &lint.RuleConfig{})
}
