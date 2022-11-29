# Advent of Code
Solutions to Advent of Code, year by year

![Christmas image of a crib with a gopher](./image.png)

## This year - 2022

I will use golang, as it is the language I use the most these days.

### TIL

#### Embed

Allows to include files and directories in a binary. Paths are relative to the directory containing the Go source file.

It allows loading a file content at runtime, in one line 🤯

Given the following file system

```txt
.
├── go.mod
├── go.sum
├── template
│   └── input.txt
|   └── template.go
└── main.go
```

```go
package template

import (
  _ "embed" // if you don't use any `embed` functions you need to import it only for its sideeffect with `_` 
  "fmt"
)

//go:embed input.txt
var inputData string

func PrintEmbed() {
  fmt.Println(inputData)
}
```

## Past years

* [2020](2020)
* [2021](2021)