package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// handler for path - обработчик пути
	http.HandleFunc("/start", start)

	// старт сервера
	fmt.Println("Starting server at port 8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func start(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}
