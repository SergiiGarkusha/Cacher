Go Cacher helps you to manage memory cache (go get github.com/SergiiGarkusha/Cacher).

See it in action:

Example:

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
