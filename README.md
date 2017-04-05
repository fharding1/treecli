# TreeCLI
## Command Line Interface Tree (TreeCLI) is a go pkg that assists in creating a tree of possible commands for CLI applications. It is meant to be fairly minimal, if you want something more full featured I'd suggest cobra, but for small tools this one is great. Typical usage may look like this:

```go
package main

import (
  "fmt"
  ct "gitlab.com/fharding/treecli"
  "os"
)

func main() {

  // !!!! ERROR HANDLING HAS BEEN OMMITED. PLEASE DEAL w/ ERRORS PROPERLY !!!!
  // !!!! CERTAIN ERRRORS CAN BE OMMITED, BUT ONLY IF THEY WILL NOT FAIL  ....`
  // !!!! (like in the AddCommand)
  // !!!! READ THE ACTUAL GO DOC !!!!

  a := ct.NewTree(func() {
    fmt.Println("no command was entered")
  })

  a.AddCommand(func() {
    fmt.Println("path: help was triggered.")
  }, "help")

 a.AddCommand(func() {
    fmt.Println("path: add was triggered.")
  }, "add")

 a.AddCommand(func() {
    fmt.Println("path: help add was triggered.")
  }, "help", "add")

  f, _ := a.Parse(os.Args[1:]...)
  f()
}
```
