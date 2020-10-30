package utils

import (
	"log"
	"regexp"
	"strings"
)

const (
	AlphaNum = 26  // alpha num count
)

// trie node definition
type TrieNode struct {
	nodes [AlphaNum]*TrieNode
	isEnd bool
}

func filterLowerLetters(s string) string {
	s = strings.ToLower(s)
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Panic(err)
	}
	return reg.ReplaceAllString(s, "")
}

// insert s into a trie tree
func (root *TrieNode) Insert(s string) {
	s = filterLowerLetters(s)

	current := root
	for _, ele := range s {
		current.nodes[ele-'a'] = &TrieNode{}
		current = current.nodes[ele-'a']
	}
	current.isEnd = true
}

// verify s in trie tree or not
func (root *TrieNode) Exist(s string) bool {
	s = filterLowerLetters(s)
	current := root
	for _, ele := range s {
		if current.nodes[ele-'a'] == nil {
			return false
		} else {
			current = current.nodes[ele-'a']
		}
	}
	return current.isEnd
}
