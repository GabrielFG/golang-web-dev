package main

import (
	"github.com/GoesToEleven/golang-web-dev/000_temp/44_class/13_interface/memcache"
	"github.com/GoesToEleven/golang-web-dev/000_temp/44_class/13_interface/cmd"
)

func main() {
	c := &memcache.MemCache{
		M: map[string]interface{}{},
	}
	cmd.CacheUser(c, "Bob", "Hello")
	cmd.CacheUser(c, "Bob", "Goodbye")
}
