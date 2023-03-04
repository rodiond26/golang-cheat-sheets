package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// implemeting Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("length of bs:", len(bs))
	return len(bs), nil
}
