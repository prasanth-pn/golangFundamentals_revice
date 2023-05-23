package main

import "fmt"


type Trie struct{
	root *TrieNode
}
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd bool
}

func main() {
	trie:=Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd: false,
		},
	}
	trie.Insert("prasanth")
	fmt.Println(trie.Search("prasant"))
}
func (t *Trie)Insert(word string){
	node:=t.root
	for _,ch:=range word{ //ranging word 
		child,ok:=node.children[ch]
		if !ok{
			child=&TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd: false,
			}
			node.children[ch]=child
		}
		node=child
	}
	node.isEnd=true
}
func (t *Trie)Search(word string)bool{
	node:=t.root
	for  _,ch:=range word{
		child, ok:=node.children[ch]
if !ok{
	return false
}
node=child
	}
	return node.isEnd
}