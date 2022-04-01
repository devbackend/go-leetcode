package trie

type Trie struct {
	isBreakpoint bool
	nodes        map[byte]*Trie
}

func Constructor() Trie {
	return Trie{
		nodes: map[byte]*Trie{},
	}
}

func (t *Trie) Insert(word string) {
	t.ins([]byte(word))
}

func (t *Trie) Search(word string) bool {
	found, exactly := t.find([]byte(word))
	return found && exactly
}

func (t *Trie) StartsWith(prefix string) bool {
	found, _ := t.find([]byte(prefix))
	return found
}

func (t *Trie) BestRoot(word string) string {
	return string(t.root([]byte(word)))
}

func (t *Trie) ins(chars []byte) {
	char := chars[0]

	if _, ok := t.nodes[char]; !ok {
		child := Constructor()
		t.nodes[char] = &child
	}

	if len(chars) == 1 {
		t.nodes[char].isBreakpoint = true
		return
	}

	t.nodes[char].ins(chars[1:])
}

func (t Trie) find(substr []byte) (found bool, exactly bool) {
	char := substr[0]

	if _, ok := t.nodes[char]; !ok {
		return
	}

	if len(substr) == 1 {
		return true, t.nodes[char].isBreakpoint
	}

	return t.nodes[char].find(substr[1:])
}

func (t *Trie) root(letters []byte) []byte {
	head := t

	var res []byte

	for _, letter := range letters {
		if _, ok := head.nodes[letter]; ok {
			res = append(res, letter)
			head = head.nodes[letter]

			continue
		}

		break
	}

	return res
}
