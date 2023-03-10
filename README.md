# Find

[![Build & Test](https://github.com/elephantoss/go-find/actions/workflows/go.yaml/badge.svg)](https://github.com/elephantoss/go-find/actions/workflows/go.yaml)
[![Coverage Status](https://coveralls.io/repos/github/elephantoss/go-find/badge.svg?branch=main&service=github)](https://coveralls.io/github/elephantoss/go-find?branch=mainn&service=github)
![GitHub](https://img.shields.io/github/license/elephantoss/go-find?color=blue)
![GitHub last commit](https://img.shields.io/github/last-commit/elephantoss/go-find?color=yellowgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/elephantoss/go-find)](https://goreportcard.com/report/github.com/elephantoss/go-find)
[![GoDoc](https://godoc.org/github.com/elephantoss/go-find?status.png)](https://godoc.org/github.com/elephantoss/go-find)

Find is a utility library that provides a convenient way to find elements in collections.


## Why do I need it?

If you are tired of writing the same 3 lines of code every time you want to get a
value out of a collection or use a default value and you wonder why there is not
a more convenient way to do it, then this lib is for you.


## Example

Get a value out of a map:

```go
value := maps.Find(myMap, "key", "some-default-value")
```


