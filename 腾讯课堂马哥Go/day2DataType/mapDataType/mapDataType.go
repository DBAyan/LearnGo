package main

import (
	"fmt"
	"sync"
)

var m = make(map[string]int)
var kownDbsMutex = &sync.Mutex{}

func main()  {
	cacheKey := "yu123"
	fmt.Println(cacheKey)
	fmt.Println(kownDbsMutex)
	kownDbsMutex.Lock()
	fmt.Println(kownDbsMutex)
	kownDbsMutex.Unlock()
	fmt.Println(kownDbsMutex)
}
