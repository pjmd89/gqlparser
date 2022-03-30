package parser

import (
	"testing"

	"github.com/pjmd89/gqlparser/v2/ast"
	"github.com/pjmd89/gqlparser/v2/parser/testrunner"
)

func TestSchemaDocument(t *testing.T) {
	testrunner.Test(t, "schema_test.yml", func(t *testing.T, input string) testrunner.Spec {
		doc, err := ParseSchema(&ast.Source{Input: input, Name: "spec"})
		return testrunner.Spec{
			Error: err,
			AST:   ast.Dump(doc),
		}
	})
}
