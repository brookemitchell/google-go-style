# Google Go Style Guide Linter

The Go linter `google-go-style` applies basic linter rules from the [google go style guide](https://google.github.io/styleguide/go-style.html)

## Example

Go source code uses MixedCaps or mixedCaps (camel case) rather than underscores (snake case) when writing multi-word names.
`my_fn` should be named `myFn` by google go convention:

```go
package main

func my_fn(format string, args ...interface{}) {
	// ...
}
```

```console
$ go vet -vettool=$(which google-go-style) ./...
./main.go:5:1: google-go-style guide `my_fn` should be named `myFn`
