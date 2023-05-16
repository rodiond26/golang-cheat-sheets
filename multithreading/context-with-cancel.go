package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i <= 10; i++ {
		go sendData(ctx, i)
	}

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)

	fmt.Println("end")
}

func sendData(ctx context.Context, num int) {
	timer := time.NewTimer(time.Duration(num) * time.Second)
	select {
	case <-ctx.Done():
		fmt.Printf("процесс #%v отменен\n", num)
		return
	case <-timer.C:
		fmt.Printf("данные процесса #%v успешно отправлены\n", num)
	}
}
