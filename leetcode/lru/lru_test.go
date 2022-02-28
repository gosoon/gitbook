package main

import (
	"fmt"
	"testing"
)

func test_lru_cache(t *testing.T) {
	cache := NewCache(2, nil)

	cache.Add("Name1", "1")
	cache.Add("Name2", "2")

	fmt.Println(cache.Get("Name1"))
	cache.Add("Name3", "3")
	fmt.Println(cache.Get("Name2"))
}
