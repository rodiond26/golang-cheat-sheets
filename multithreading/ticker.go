package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	ticker := time.NewTicker(1 * time.Second)
	count := 0

	for tick := range ticker.C {
		count++
		fmt.Printf("Tick #%v : time %v\n", count, tick)
		if count > 7 {
			ticker.Stop()
			break
		}
	}

	fmt.Println("end")

	// тикер через канал
	fmt.Println("start")

	ticker2 := time.Tick(1 * time.Second)
	count2 := 0

	for tick := range ticker2 {
		count2++
		fmt.Printf("Tick2 #%v : time %v\n", count, tick)
		// этот тикер нельзя остановить
	}

	fmt.Println("end")
}
