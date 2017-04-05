package clitree

import (
  "fmt"
)

// CLITree ...
// Root tree for commands.
type CLITree struct {
  root *Node
}

// Node ...
// Holds a command node.
type Node struct {
  f        func()
  children map[string]*Node
}

// NewTree ...
// Creates a new tree. The given func
// should be the function to give if no params
// are given. (help)
func NewTree(f func()) *CLITree {
  return &CLITree{&Node{f, make(map[string]*Node)}}
}

// AddCommand ...
// Adds a command to the CLI tree.
func (ct *CLITree) AddCommand(f func(), path ...string) error {
  cur := ct.root
  for i, v := range path {
    if i == len(path)-1 && cur.children[v] == nil {
      cur.children[v] = &Node{f, make(map[string]*Node)}
    } else if _, ok := cur.children[v]; ok == true {
      cur = cur.children[v]
    } else {
      return fmt.Errorf("error: does not contain node %s in path at position %d of %d", v, i, len(path)-1)
    }
  }
  return nil
}

// RemoveCommand ...
// Removes a command from the CLI tree.
func (ct *CLITree) RemoveCommand(path ...string) error {
  n, k, err := ct.searchTree(path...)
  if err != nil {
    return err
  }
  n.children[k] = nil
  return nil
}

// Print ...
// Kinda pretty prints the CLI tree.
func (ct *CLITree) Print() {
  var printChildren func(*Node, int)
  printChildren = func(n *Node, indent int) {
    for k, v := range n.children {
      fmt.Printf("%*s%s\n", indent, "", k)
      printChildren(v, indent+2)
    }
  }
  printChildren(ct.root, 0)
}

// Parse ...
// Parses a path (variadic strings) and returns
// the function present from that path.
func (ct *CLITree) Parse(path ...string) (func(), error) {
  if len(path) == 0 {
    return ct.root.f, nil
  }
  n, k, err := ct.searchTree(path...)
  if err != nil {
    return nil, err
  }
  return n.children[k].f, nil
}

func (ct *CLITree) searchTree(path ...string) (*Node, string, error) {
  cur := ct.root
  for i, v := range path {
    if i == len(path)-1 && cur.children[v] != nil {
      return cur, v, nil
    } else if i == len(path)-1 && cur.children[v] == nil {
      return nil, "", fmt.Errorf("error: that node does not exist")
    } else if _, ok := cur.children[v]; ok == true {
      cur = cur.children[v]
    } else {
      fmt.Println(i, v, path)
      return nil, "", fmt.Errorf("error: does not contain node %s where path does", v)
    }
  }
  return nil, "", fmt.Errorf("error: tree does not contain node in path")
}
