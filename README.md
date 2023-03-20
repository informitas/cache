Go Cache
================================

Go Cache helps you to cache your data in a simple way.

## Installation

```go
go get github.com/informitas/cache
```

## Usage Basic

```go
package main

import (
	"fmt"

	"github.com/informitas/cache"
)

func main() {
	cache := cache.New[string]()
	cache.Set("userID", "12345");

	userID, err := cache.Get("userID")
	if err != nil {
		panic(err)
	}
	fmt.Println(userID) // 12345

	cache.Set("first", "1");
	cache.Set("second", "2");

	fmt.Println(cache.Keys()) // [first second]

	fmt.Println(cache.Size()) // [2] size of cache

	cache.Clear()
	fmt.Println(cache.Keys()) // []
}


```


## Usage With TTL Options
```go
package main

import (
	"fmt"
	"time"

	"github.com/informitas/cache"
)

func main() {
	cache := cache.New[int]()

	options := cache.Options()
	options.TTL = 2 * time.Second

	err := cache.Set("key", 55, options)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(3 * time.Second)

	value, err := cache.Get("key")
	if err != nil {
		fmt.Println(err) // err because key is expired
	}
	fmt.Println(value)
}

```

## Usage With Immutable Options
```go
package main

import (
	"fmt"

	"github.com/informitas/cache"
)

func main() {
	cache := cache.New[int]()

	options := cache.Options()
	options.Immutable = true

	err := cache.Set("key", 55, options)
	if err != nil {
		fmt.Println(err)
	}

	err = cache.Set("key", 100)
	if err != nil {
		fmt.Println(err) // err because value is immutable
	}
}


```
