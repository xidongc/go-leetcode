package main

// 212 word search II: using dfs with trie to prune, using hash set to de-dup
func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return []string {}
	}
	// create trie
	root := CreateTrie()
	for _, word := range words {
		root.Insert(word)
	}

	dict := map[string]struct{}{}

	for _, word := range words {
		dict[word] = struct{}{}
	}

	results := map[string]struct{}{}
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[0]); j ++ {
			visited := map[[2]int]struct{}{
				[2]int{i, j}: {},
			}
			findWordsHelper(root, board, dict, i, j, results, []byte{board[i][j]}, visited)
		}
	}

	arrResult := make([]string, 0)
	for key, _ := range results {
		arrResult = append(arrResult, key)
	}
	return arrResult
}

// dfs helper
func findWordsHelper(trie *TrieNode,board [][]byte, words map[string]struct{}, x int, y int, results map[string]struct{}, tmp []byte, visited map[[2]int]struct{}) {
	if !trie.Search(string(tmp)) {
		return
	}
	if _, exist := words[string(tmp)]; exist {
		results[string(tmp)] = struct{}{}
	}

	posArray := [4][2]int{{0, 1},{0, -1}, {-1, 0}, {1, 0}}
	for _, pos := range posArray {
		if inBound(x+pos[0], y+pos[1], len(board), len(board[0])) {
			if _, exist := visited[[2]int{x+pos[0],y+pos[1]}]; exist {
				continue
			}
			visited[[2]int{x+pos[0],y+pos[1]}] = struct{}{}
			tmp = append(tmp, board[x+pos[0]][y+pos[1]])
			findWordsHelper(trie, board, words, x+pos[0], y+pos[1], results, tmp, visited)
			tmp = tmp[0 : len(tmp)-1]
			delete(visited, [2]int{x+pos[0],y+pos[1]})
		}
	}
}

const AlphCount = 26

type TrieNode struct {
	nodes [AlphCount]*TrieNode
	isEnd bool
}

func CreateTrie() *TrieNode {
	return &TrieNode{}
}

func (root *TrieNode)Insert(word string) {
	current := root
	for _, w := range word {
		if current.nodes[w-'a'] == nil {
			current.nodes[w-'a'] = &TrieNode{}
		}
		current = current.nodes[w-'a']
	}
	current.isEnd = true
}

// determine prefix
func (root *TrieNode)Search(word string) bool {
	current := root
	for _, w := range word {
		if current.nodes[w-'a'] == nil {
			return false
		}
		current = current.nodes[w-'a']
	}
	return true
}

// determine in bound or not
func inBound(x int, y int, length int, width int) bool {
	if 0 <= x && x < length && y >= 0 && y < width {
		return true
	}
	return false
}
