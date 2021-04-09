package main

import (
	"container/list"
)

/*
	127 word ladder I
	A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words such that:

	The first word in the sequence is beginWord.
	The last word in the sequence is endWord.
	Only one letter is different between each adjacent pair of words in the sequence.
	Every word in the sequence is in wordList.
	Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in the shortest transformation
	sequence from beginWord to endWord, or 0 if no such sequence exists.

	basically it is shortest path in a graph, use bfs level traversal
	iterate word with one letter differ from a-z, and determine if it is in wordDict
	turn wordDict into hashSet first, for O(1) find

Example 1:
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog" with 5 words.

Example 2:
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: The endWord "cog" is not in wordList, therefore there is no possible transformation.
*/
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

/*
	126 WordLadderII
	Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation
	sequence(s) from beginWord to endWord, such that:

	Only one letter can be changed at a time
	Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
	Note:

	Return an empty list if there is no such transformation sequence.
	All words have the same length.
	All words contain only lowercase alphabetic characters.
	You may assume no duplicates in the word list.
	You may assume beginWord and endWord are non-empty and are not the same.

	using bfs to find depth, and get node -> depth relationship, using dfs to recursively
	find all paths with the rel got from bfs result

Example 1:
Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]

Example 2:
Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: []

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.
*/
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
