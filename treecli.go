package treecli

import (
	"fmt"
)

// TreeCLI holds the root node for all nodes to stem off.
type TreeCLI struct {
	root *Node
}

// Node holds a command node with a function that is returned from looking up a path,
// for example if the path is "auth", children might be ["token", "user"] for the total
// path "auth token" or "auth user".
type Node struct {
	f        func()
	children map[string]*Node
}

// NewTree creates a new tree. The given function will be returned if no params are given
// when parsing a path.
func NewTree(f func()) *TreeCLI {
	return &TreeCLI{&Node{f, make(map[string]*Node)}}
}

// AddCommand adds a command node to the tree, given a function to associate with the
// specified path.
func (ct *TreeCLI) AddCommand(f func(), path ...string) error {
	cur := ct.root
	for i, v := range path {
		if v == "*" && i != len(path)-1 {
			return fmt.Errorf("error: wildcards cannot have children")
		} else if i == len(path)-1 && cur.children[v] == nil {
			cur.children[v] = &Node{f, make(map[string]*Node)}
		} else if _, ok := cur.children[v]; ok == true {
			cur = cur.children[v]
		} else {
			return fmt.Errorf("error: does not contain node %s in path at position %d of %d", v, i, len(path)-1)
		}
	}
	return nil
}

// RemoveCommand removes a command from the tree, given a path to that node.
func (ct *TreeCLI) RemoveCommand(path ...string) error {
	n, k, err := ct.searchTree(path...)
	if err != nil {
		return err
	}
	n.children[k] = nil
	return nil
}

// String pretty prints the tree.
func (ct *TreeCLI) String() (str string) {
	var printChildren func(*Node, int)
	printChildren = func(n *Node, indent int) {
		for k, v := range n.children {
			str += fmt.Sprintf("%*s%s\n", indent, "", k)
			printChildren(v, indent+2)
		}
	}
	printChildren(ct.root, 0)
	return str
}

// Parse parses a path (variadic strings) and returns the function at that node.
func (ct *TreeCLI) Parse(path ...string) (func(), error) {
	if len(path) == 0 {
		return ct.root.f, nil
	}
	n, k, err := ct.searchTree(path...)
	if err != nil {
		return nil, err
	}
	return n.children[k].f, nil
}

func (ct *TreeCLI) searchTree(path ...string) (*Node, string, error) {
	cur := ct.root
	for i, v := range path {
		if cur.children["*"] != nil {
			return cur, "*", nil
		} else if i == len(path)-1 && cur.children[v] != nil {
			return cur, v, nil
		} else if i == len(path)-1 && cur.children[v] == nil {
			return nil, "", fmt.Errorf("error: that node does not exist")
		} else if _, ok := cur.children[v]; ok == true {
			cur = cur.children[v]
		} else {
			return nil, "", fmt.Errorf("error: does not contain node %s where path does", v)
		}
	}
	return nil, "", fmt.Errorf("error: tree does not contain node in path")
}
