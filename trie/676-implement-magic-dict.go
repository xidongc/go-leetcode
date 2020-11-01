package main

const AlphaCount = 26

type MagicDictionary struct {
	nodes [AlphaCount]*MagicDictionary
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (dict *MagicDictionary) BuildDict(dictionary []string)  {
	for _, str := range dictionary {
		dict.constructMagic(str)
	}
}

func (dict *MagicDictionary) constructMagic(str string) {
	current := dict
	for i := 0; i < len(str); i ++ {
		// previously forget to verified == situation led to failed
		if current.nodes[str[i]-'a'] == nil {
			current.nodes[str[i]-'a'] = &MagicDictionary{}
		}
		current = current.nodes[str[i] - 'a']
	}
	current.isEnd = true
}

func (dict *MagicDictionary) Search(searchWord string) bool {
	return dict.searchWithTolerance(dict, searchWord, 1)
}

func (dict *MagicDictionary) searchWithTolerance(current *MagicDictionary, searchWord string, tolerance int) bool {
	if len(searchWord) == 0 {
		return current.isEnd && tolerance == 0
	}
	if tolerance < 0 {
		return false
	}

	verified := false
	for j := 0; j < AlphaCount; j ++ {
		if current.nodes[j] != nil {
			if int(searchWord[0]-'a') != j {
				verified = verified || dict.searchWithTolerance(current.nodes[j], searchWord[1:], tolerance-1)
			} else {
				verified = verified || dict.searchWithTolerance(current.nodes[j], searchWord[1:], tolerance)
			}
		}
	}
	return verified
}
/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
