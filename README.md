Go Cache
================================

Go Cache helps you to cache your data in a simple way.

## Installation

```go
go get github.com/informitas/cache
```

See it in action:

```go
package main

import (
	"fmt"

	"github.com/informitas/cache"
)

func main() {
	cache := cache.New();
	cache.Set("userID", "12345");

	userID := cache.Get("userID").(string)
	fmt.Println(userID) // 12345

	cache.Delete("userID")
	if cache.Has("userID") {
		fmt.Print("userID: %s", cache.Get("userID").(string))
	} else {
		fmt.Println("userID is not in cache")
	}

	cache.Set("first", "1");
	cache.Set("second", "2");

	fmt.Println(cache.Keys()) // [first second]

	fmt.Println(cache.Size()) // [2] size of cache

	cache.Clear()
	fmt.Println(cache.Keys()) // []
}

```
