# argoments [![Go Reference](https://pkg.go.dev/badge/github.com/elymination/argoments.svg)](https://pkg.go.dev/github.com/elymination/argoments) 

Go package to register and use command-line arguments in your Go program.

## How to use
Get the package with `go get github.com/elymination/argoments` and import it to your project.

**Example:**
```
package main

import (
    "fmt"
    "github.com/elymination/argoments"
)

func main() {
    args := argoments.Init()
    args.RegisterParamed([]string{"--size", "-r", "--repeat"})
    args.Parse()
    sizeArgValue, err := args.GetValue("size")
    if err != nil {
        fmt.Println(sizeArgValue)
    }
}
```
