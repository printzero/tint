### Getting Started Guide

Tint has some really great functions to get you started with coloring your terminal output. We first need to import
tint and initialize a pointer to tint. For that we have a `Init()` function.

Example:

```go
package main
import "github.com/printzero/tint"

func main() {
	t := tint.Init()
}
```

The next steps are written assuming you have initialized tint like the above example.

> **t.Raw()**

This functions lets you color a string and returns a colored string.



