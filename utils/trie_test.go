package utils

import (
	"testing"
)

func TestTrie(t *testing.T) {
	root := TrieNode{}
	root.Insert("hello world")
	if root.Exist("hello world") != true {
		t.Error("exist failed")
	}
	if root.Exist("hello wol") != false {
		t.Error("not exist failed")
	}
}
