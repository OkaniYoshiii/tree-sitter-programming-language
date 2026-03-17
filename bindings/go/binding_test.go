package tree_sitter_pl_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_pl "github.com/tree-sitter/tree-sitter-pl/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_pl.Language())
	if language == nil {
		t.Errorf("Error loading Programming Language grammar")
	}
}
