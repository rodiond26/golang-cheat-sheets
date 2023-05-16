package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("start")
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go sleep(1*time.Second, wg)

	wg.Add(1)
	go sleep(2*time.Second, wg)

	wg.Add(1)
	go sleep(3*time.Second, wg)

	wg.Wait()
	fmt.Println("end")
}

func sleep(t time.Duration, wg *sync.WaitGroup) {
	time.Sleep(t)
	fmt.Println("sleep", t)
	wg.Done()
}
