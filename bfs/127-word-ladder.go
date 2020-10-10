package main

import (
	"container/list"
)

// 127 word ladder I, use bfs level traversal as we only care about shortest path number
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}
	if beginWord == endWord {
		return 1
	}

	transitionNum := 1
	hashset := map[string]struct{}{}

	queue := list.New()
	queue.PushBack(beginWord)
	hashset[beginWord] = struct{}{}

	var dict = make(map[string]struct{}, 0)
	for _, dictWord := range wordList {
		dict[dictWord] = struct {}{}
	}

	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			ele := queue.Front().Value.(string)
			queue.Remove(queue.Front())
			for _, word := range nextWord(dict, ele) {
				if word == endWord {
					return transitionNum + 1
				}
				if _, exist := hashset[word]; !exist {
					queue.PushBack(word)
					hashset[word] = struct{}{}
				}
			}
		}
		transitionNum += 1
	}
	return 0
}

func nextWord(dict map[string]struct{}, word string) []string {
	nextWordList := make([]string, 0)
	for i := 0; i < len(word); i++ {
		// loop 26 letters
		for j := 0; j < 26; j++ {
			if word[i] == byte(int('a')+j) {
				continue
			}
			tmp := word[:i] + string(int('a')+j) + word[i+1:]
			if containsStr(dict, tmp) {
				nextWordList = append(nextWordList, tmp)
			}
		}
	}
	return nextWordList
}

func containsStr(dict map[string]struct{}, target string) bool {
	_, exist := dict[target]
	return exist
}

// 126 WordLadderII bfs to find depth, dfs to recursively find all paths
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var dict = make(map[string]struct{}, 0)
	for _, dictWord := range wordList {
		dict[dictWord] = struct {}{}
	}
	depth := markDepth(beginWord, endWord, dict)
	result := make([][]string, 0)
	wordLadderHelper(depth, &result, beginWord, beginWord, endWord, dict, []string{beginWord})
	return result
}

// bfs
func markDepth(beginWord string, endWord string, dict map[string]struct{}) map[string]int {
	queue := list.New()
	depth := make(map[string]int)

	queue.PushBack(beginWord)
	depth[beginWord] = 1

	for queue.Len() > 0 {
		ele := queue.Front().Value.(string)
		queue.Remove(queue.Front())

		for _, word := range nextWord(dict, ele) {
			if _, exist := depth[word]; !exist {
				depth[word] = depth[ele] + 1
				queue.PushBack(word)
				if word == endWord {
					return depth
				}
			}
		}
	}
	return nil
}

// dfs
func wordLadderHelper(
	depth map[string]int, results *[][]string,
	current string, startWord string,
	endWord string, dict map[string]struct{},
	result []string) {

	if result[len(result)-1] == endWord {
		tmp := make([]string, len(result), len(result))
		copy(tmp, result)
		*results = append(*results, tmp)
	}

	for _, word := range nextWord(dict, current) {
		if depth[word] != depth[current] + 1 {
			continue
		}
		result = append(result, word)
		wordLadderHelper(depth, results, word, startWord, endWord, dict, result)
		result = result[:len(result)-1]
	}
}
