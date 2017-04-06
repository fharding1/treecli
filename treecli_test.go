package treecli

import "testing"

// TestBasicFunctionality ...
// Tests basic functionality (creating tree),
// having it modify a var, calling a func,
// using parse, etc. function...
func TestBasicFunctionality(t *testing.T) {
  res := ""
  ct := NewTree(func() {
    res = "test"
  })
  f, err := ct.Parse()
  if err != nil {
    t.Error("error: had error when it should not: " + err.Error())
  }
  f()
  if res != "test" {
    t.Error("error: did not execute function properly from correct path and setup")
  }
}

// TestAddCommandAndParse ...
func TestAddCommandAndParse(t *testing.T) {
  res := ""
  ct := NewTree(func() {})
  if err := ct.AddCommand(func() { res = "test" }, "test"); err != nil {
    t.Error("error: nothing was wrong, should have added")
  }
  if err := ct.AddCommand(func() { /* do nothing */ }, "a", "b", "c"); err == nil {
    t.Error("error: should have failed, path not created yet")
  }
  if err := ct.AddCommand(func() { res = "test2" }, "test", "2"); err != nil {
    t.Error("error: nothing was wrong, should have added")
  }
  if err := ct.AddCommand(func() { /* do nothing */ }, "test", "*"); err != nil {
    t.Error("error: nothing was wrong, should have added")
  }
  if err := ct.AddCommand(func() { /* do nothing */ }, "test", "*", "*"); err == nil {
    t.Error("error: added child to wildcard")
  }

  f, err := ct.Parse("test")
  if err != nil {
    t.Error("error: nothing was wrong, should have parsed")
  }
  f()
  if res != "test" {
    t.Error("error: function should have changed res to 'test'")
  }
  f, err = ct.Parse("test", "2")
  if err != nil {
    t.Error("error: nothing was wrong, should have parsed")
  }
}

// TestRemoveCommand ...
func TestRemoveCommand(t *testing.T) {
  ct := NewTree(func() {})
  ct.AddCommand(func() {}, "a")
  _, err := ct.Parse("a")
  if err != nil {
    t.Error("error: adding command so that it can be removed")
  }
  err = ct.RemoveCommand("a")
  if err != nil {
    t.Error("error: should have removed successfully")
  }

  err = ct.RemoveCommand("a")
  if err == nil {
    t.Error("error: shouldn't have removed successfully.")
  }
}
