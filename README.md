# hackernews-api-go

wrapper for hacker news API ( https://github.com/HackerNews/API ) written in Golang.


## Sample Usage

```
package main

import (
	h "github.com/qube81/hackernews-api-go"

	"fmt"
)

func main() {

	a, _ := h.GetItem(8863)
	fmt.Printf("%+v", a)

	b, _ := h.GetStories("ask")
	fmt.Printf("%+v", b)

	c, _ := h.GetMaxItemID()
	fmt.Printf("%+v", c)

	d, _ := h.GetUser("jl")
	fmt.Printf("%+v", d)

	e, _ := h.GetUpdates()
	fmt.Printf("%+v", e)

}
```