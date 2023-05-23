package main

import "fmt"

type Tree struct {
	Root *Node
}
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func (t *Tree) Insert(data int) {
	newnode := new(Node)
	newnode.Data, newnode.Left, newnode.Right = data, nil, nil
	if t.Root == nil {
		t.Root = newnode
	} else {
		current := t.Root
		for current != nil {
			if data < current.Data {
				if current.Left == nil {
					current.Left = newnode
					return
				}
				current = current.Left
			} else {
				if current.Right == nil {
					current.Right = newnode
					return
				}
				current = current.Right
			}
		}
	}

}
func main() {
	tree := Tree{}
	a := []int{23, 45, 12, 56, 34, 11}
	for _, v := range a {
		tree.Insert(v)
	}

	tree.Search(45)
	//tree.Delete(12)
	inorder(tree.Root)
	fmt.Println("\n--------------------------------------------------")
	postorder(tree.Root)
	fmt.Println("\n----------------------------------------------------")
	preorder(tree.Root)
	fmt.Println("\n------------------------------------------------------\n ")

}
func inorder(root *Node) {
	if root != nil {
		inorder(root.Left)
		fmt.Print("\t", root.Data)
		inorder(root.Right)
	}

}
func postorder(root *Node) {
	if root != nil {
		postorder(root.Left)
		postorder(root.Right)
		fmt.Print("\t", root.Data)
	}
}
func preorder(root *Node) {
	if root != nil {
		fmt.Print("\t", root.Data)
		preorder(root.Left)
		preorder(root.Right)
	}
}

// searching in tree
func (tree *Tree) Search(data int) {
	if tree.Root.Data == data {
		fmt.Println("data is available")
	} else {
		current := tree.Root
		for current != nil {
			if data < current.Data {
				current = current.Left
			} else if data > current.Data {
				current = current.Right

			} else {
				fmt.Println("the data is present")
				return
			}
		}
		fmt.Println("data is not present")
	}
}
func (t *Tree) Delete(data int) {
	t.Remove(data, t.Root, nil)

}
func (t *Tree) Remove(data int, currentNode, parentNode *Node) {

	for currentNode != nil {
		if data < currentNode.Data {
			parentNode = currentNode
			currentNode = currentNode.Left
		} else if data > currentNode.Data {
			parentNode = currentNode
			currentNode = currentNode.Right
		} else {
			if currentNode.Left != nil && currentNode.Right != nil {
				currentNode.Data = getMinimum(currentNode.Right)
				t.Remove(currentNode.Data, currentNode.Right, currentNode)
			} else {
				if parentNode == nil {
					if currentNode.Left == nil {
						t.Root = currentNode.Right
					} else {
						t.Root = currentNode.Left
					}
				} else {
					if parentNode.Left == currentNode {
						if parentNode.Right == nil {
							parentNode.Left = currentNode.Left
						} else {
							parentNode.Left = currentNode.Left
						}
					} else {
						if parentNode.Right == nil {
							parentNode.Right = currentNode.Left

						} else {
							parentNode.Right = currentNode.Right
						}
					}
				}
			}
			break
		}

	}

}
func getMinimum(currentNode *Node) int {
	if currentNode.Left == nil {
		return currentNode.Data
	} else {
		return getMinimum(currentNode.Left)
	}
}
