package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	parentCtx := context.Background()

	ctx, _ := context.WithTimeout(parentCtx, 2*time.Second)
	for i := 1; i <= 10; i++ {
		go sendData2(ctx, i)
	}

	time.Sleep(3 * time.Second)
	fmt.Println("end")
}

func sendData2(ctx context.Context, num int) {
	timer := time.NewTimer(time.Duration(num) * time.Second)
	select {
	case <-ctx.Done():
		fmt.Printf("процесс #%v отменен\n", num)
		return
	case <-timer.C:
		fmt.Printf("данные процесса #%v успешно отправлены\n", num)
	}
}
