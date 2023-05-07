package main

import (
	"fmt"
)

func main() {
	services := GenerateCheck()
	for i, _ := range services {
		if services[i].Status == PassStatus {
			fmt.Println(services[i].ServiceID)
		}
	}
}
