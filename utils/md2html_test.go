package utils

import (
	"fmt"
	"testing"
)

func TestMarkdownToHTML(t *testing.T) {
	path := "../test/test.md"
	res, err := ReadAll(path)
	if err != nil {
		t.Error(fmt.Sprintf("read file %v fail", err))
	}

	result := MarkdownToHTML(string(res))
	t.Log(result)
}
