package trie

// https://w3.cs.jmu.edu/lam2mo/cs240_2014_08/pa05-tries.html

type TrieNode struct {
	isEndOfWord bool
	children    map[rune]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		isEndOfWord: false,
		children:    make(map[rune]*TrieNode),
	}
}

// Insert a word into the trie
func (t *TrieNode) Insert(word string) {
	current := t

	for _, char := range word {
		child, exists := current.children[char]
		if !exists {
			child = NewTrieNode()
			current.children[char] = child
		}
		current = child
	}

	// check if the word is already in the trie
	if current.isEndOfWord {
		return
	}

	current.isEndOfWord = true
}

// Search a word in the trie
func (t *TrieNode) Search(word string) bool {
	current := t

	for _, char := range word {
		child, exists := current.children[char]
		if !exists {
			return false
		}
		current = child
	}

	return current.isEndOfWord
}

// Check if the trie contains prefix
func (t *TrieNode) StartsWith(prefix string) bool {
	current := t

	for _, char := range prefix {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}

	return true
}

// Spellchecker with suggestions
func (t *TrieNode) SpellCheck(word string) []string {
	return nil
}

// TODO:
// - simple spellchecker with suggestions
