package treecli

import (
  "fmt"
)

// TreeCLI ...
// Root tree for your command tree.
// Simply holds a root node for all
// to stem off of.
type TreeCLI struct {
  root *Node
}

// Node ...
// Holds a command node with a function
// that is returned when looking up a path,
// and children nodes that are children of that
// path. For example if this is path "auth",
// children might be ["token", "user"] for the
// total path "auth token" or "auth user".
type Node struct {
  f        func()
  children map[string]*Node
}

// NewTree ...
// Creates a new tree. The given func
// should be the function to give if no params
// are given. (help)
func NewTree(f func()) *TreeCLI {
  return &TreeCLI{&Node{f, make(map[string]*Node)}}
}

// AddCommand ...
// Adds a command node to the CLI tree. Given
// a function to associate w/ the specified
// path.
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

// SimpleMessage ...
// This is a helper function for making a command
// node that simply says a message, for example for
// help messages.
func (ct *TreeCLI) SimpleMessage(msg string, path ...string) error {
  cur := ct.root
  for i, v := range path {
    if v == "*" && i != len(path)-1 {
      return fmt.Errorf("error: wildcards cannot have children")
    } else if i == len(path)-1 && cur.children[v] == nil {
      cur.children[v] = &Node{func() {
        fmt.Println(msg)
      }, make(map[string]*Node)}
    } else if _, ok := cur.children[v]; ok == true {
      cur = cur.children[v]
    } else {
      return fmt.Errorf("error: does not contain node %s in path at position %d of %d", v, i, len(path)-1)
    }
  }
  return nil
}

// RemoveCommand ...
// Removes a command from the CLI tree. Given
// the path to that command node.
func (ct *TreeCLI) RemoveCommand(path ...string) error {
  n, k, err := ct.searchTree(path...)
  if err != nil {
    return err
  }
  n.children[k] = nil
  return nil
}

// String ...
// Kinda pretty prints the CLI tree recursively.
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

// Parse ...
// Parses a path (variadic strings) and returns
// the function present from that path.
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
