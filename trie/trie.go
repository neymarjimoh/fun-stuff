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
	var suggestions []string

	// helper function to perform DFS on the trie
	var dfs func(node *TrieNode, prefix string)
	dfs = func(node *TrieNode, prefix string) {
		if len(suggestions) >= 5 {
			return // no more suggestions with 5 as limit
		}

		if node == nil {
			return
		}

		if node.isEndOfWord {
			suggestions = append(suggestions, prefix)
		}

		for char, child := range node.children {
			dfs(child, prefix+string(char))
		}
	}

	current := t
	for i := 0; i < len(word); i++ {
		char := rune(word[i])
		child, exists := current.children[char]
		if !exists {
			return suggestions // no suggestions if word not found
		}
		current = child
	}

	// If the word is found, check suggestions for words that can be formed by adding one character
	for ch, child := range current.children {
		dfs(child, word+string(ch))
	}

	return suggestions
}
