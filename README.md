Go Cacher helps you to manage memory cache.

See it in action:

package main

import (
	"fmt"
	"inmemorycache/cache"
)

func main() {

	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")
	fmt.Println(userId)
	fmt.Println(cache)

	cache.Set("userId1", 15)

	userId = cache.Get("userId1")
	fmt.Println(userId)
	fmt.Println(cache)

	cache.Delete("userId")
	fmt.Println(cache)

	userId = cache.Get("userId1")
	fmt.Println(userId)

}
