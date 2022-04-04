package medium

import (
	"strings"
)

// ReplaceWords for https://leetcode.com/problems/replace-words/
func ReplaceWords(dictionary []string, sentence string) string {
	tr := newTrie()

	for i := range dictionary {
		tr.insert([]byte(dictionary[i]))
	}

	words := strings.Split(sentence, " ")

	for k := range words {
		root := tr.root([]byte(words[k]))
		if len(root) != 0 {
			words[k] = string(root)
		}
	}

	return strings.Join(words, " ")
}

type trie struct {
	isBreakpoint bool
	nodes        map[byte]*trie
}

func newTrie() trie {
	return trie{
		nodes: map[byte]*trie{},
	}
}

func (t *trie) insert(chars []byte) {
	char := chars[0]

	if _, ok := t.nodes[char]; !ok {
		child := newTrie()
		t.nodes[char] = &child
	}

	if len(chars) == 1 {
		t.nodes[char].isBreakpoint = true
		return
	}

	t.nodes[char].insert(chars[1:])
}

func (t *trie) root(letters []byte) []byte {
	head := t

	var res []byte

	for _, letter := range letters {
		if _, ok := head.nodes[letter]; ok {
			res = append(res, letter)
			head = head.nodes[letter]

			if head.isBreakpoint {
				return res
			}

			continue
		}

		break
	}

	return []byte{}
}
