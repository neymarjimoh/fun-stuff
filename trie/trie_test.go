package trie

import (
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := NewTrieNode()

	trie.Insert("apple")
	trie.Insert("banana")
	trie.Insert("cherry")

	// Assert the words are present in the Trie
	assertSearch(t, trie, "apple", true)
	assertSearch(t, trie, "banana", true)
	assertSearch(t, trie, "cherry", true)

	// Assert other words are not present
	assertSearch(t, trie, "grape", false)
	assertSearch(t, trie, "melon", false)
}

func TestTrie_Search(t *testing.T) {
	trie := NewTrieNode()

	trie.Insert("apple")
	trie.Insert("banana")
	trie.Insert("cherry")

	// Assert the search results
	assertSearch(t, trie, "apple", true)
	assertSearch(t, trie, "banana", true)
	assertSearch(t, trie, "cherry", true)

	// Assert other words are not found
	assertSearch(t, trie, "grape", false)
	assertSearch(t, trie, "melon", false)
}

func TestTrie_StartsWith(t *testing.T) {
	trie := NewTrieNode()

	trie.Insert("apple")
	trie.Insert("banana")
	trie.Insert("cherry")

	// Assert prefix matches
	assertStartsWith(t, trie, "app", true)
	assertStartsWith(t, trie, "ban", true)
	assertStartsWith(t, trie, "che", true)

	// Assert prefix doesn't match
	assertStartsWith(t, trie, "gra", false)
	assertStartsWith(t, trie, "mel", false)
}

func assertSearch(t *testing.T, trie *TrieNode, word string, expected bool) {
	t.Helper()

	if trie.Search(word) != expected {
		t.Errorf("Search failed for word '%s'. Expected: %v, Got: %v", word, expected, !expected)
	}
}

func assertStartsWith(t *testing.T, trie *TrieNode, prefix string, expected bool) {
	t.Helper()

	if trie.StartsWith(prefix) != expected {
		t.Errorf("StartsWith failed for prefix '%s'. Expected: %v, Got: %v", prefix, expected, !expected)
	}
}

func TestTrie_SpellCheck(t *testing.T) {
	trie := NewTrieNode()

	trie.Insert("apple")
	trie.Insert("apply")
	trie.Insert("banana")
	trie.Insert("cherry")
	trie.Insert("grape")
	trie.Insert("melon")

	misspelled := "app"
	expectedSuggestions := []string{"apple", "apply"}
	suggestions := trie.SpellCheck(misspelled)

	// Assert the suggestions
	if len(suggestions) != len(expectedSuggestions) {
		t.Errorf("SpellCheck failed for word '%s'. Expected suggestions: %v, Got: %v", misspelled, expectedSuggestions, suggestions)
	} else {
		for i := 0; i < len(suggestions); i++ {
			if suggestions[i] != expectedSuggestions[i] {
				t.Errorf("SpellCheck failed for word '%s'. Expected suggestions: %v, Got: %v", misspelled, expectedSuggestions, suggestions)
				break
			}
		}
	}
}
