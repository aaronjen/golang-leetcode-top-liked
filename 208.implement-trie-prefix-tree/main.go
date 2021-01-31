package main

/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
type node struct {
	isWord   bool
	children []*node
}

// Trie struct
type Trie struct {
	root *node
}

// Constructor func
/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: &node{
			children: make([]*node, 26),
		},
	}
}

// Insert func
/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	nd := t.root
	for i := 0; i < len(word); i++ {
		pos := int(word[i] - 'a')
		if nd.children[pos] == nil {
			nd.children[pos] = &node{
				children: make([]*node, 26),
			}
		}
		nd = nd.children[pos]
	}
	nd.isWord = true
}

// Search func
/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	nd := t.root
	for i := 0; i < len(word); i++ {
		pos := int(word[i] - 'a')
		if nd.children[pos] == nil {
			return false
		}
		nd = nd.children[pos]
	}
	return nd.isWord
}

// StartsWith func
/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	nd := t.root
	for i := 0; i < len(prefix); i++ {
		pos := int(prefix[i] - 'a')
		if nd.children[pos] == nil {
			return false
		}
		nd = nd.children[pos]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
