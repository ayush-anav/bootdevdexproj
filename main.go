package main

import (
	"ayush-anav/bootdevdexproj/internal/pokecache"
	"time"
)

func main() {
	c := &config{
		cache: pokecache.NewCache(5 * time.Second),
	}
	startRepl(c)
}
