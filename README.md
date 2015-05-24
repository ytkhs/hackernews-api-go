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

	a, _ := h.GetItem(9588901)
	fmt.Println(a)

	b, _ := h.GetStories("ask")
	fmt.Println(b)

	c, _ := h.GetMaxItemID()
	fmt.Println(c)

	d, _ := h.GetUser("jl")
	fmt.Println(d)

	e, _ := h.GetUpdates()
	fmt.Println(e)

}
```