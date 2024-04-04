# go-peek

Package **peek** provides tools for peeking at input, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-peek

[![GoDoc](https://godoc.org/github.com/reiver/go-peek?status.svg)](https://godoc.org/github.com/reiver/go-peek)

## Example

Here is an example:

```golang
import "github.com/reiver/go-peek"

// ...

r, size, err := peek.PeekRune(runescanner)
```

## Import

To import package **peek** use `import` code like the follownig:
```
import "github.com/reiver/go-peek"
```

## Installation

To install package **peek** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-peek
```

## Author

Package **peek** was written by [Charles Iliya Krempeaux](http://reiver.link)
