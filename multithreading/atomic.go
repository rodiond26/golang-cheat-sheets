package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	cl := &Client2{}

	for i := 1; i <= 100000; i++ {
		go addAge2(cl, 1)
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("%#v\n", cl)
}

func addAge2(cl *Client2, add int) {
	atomic.AddInt64(&cl.Age, int64(add))
}

type Client2 struct {
	Age int64
}
