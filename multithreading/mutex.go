package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cl := &Client{}
	mu := &sync.Mutex{}

	for i := 1; i <= 100000; i++ {
		go addAge(cl, 1, mu)
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("%#v\n", cl)
}

func addAge(cl *Client, add int, mu *sync.Mutex) {
	mu.Lock()
	cl.Age = cl.Age + add
	mu.Unlock()
}

type Client struct {
	Age int
}
