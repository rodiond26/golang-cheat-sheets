package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Response: %+v\n", resp)
	fmt.Printf("Response body: %+v\n", resp.Body)

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Printf(string(bs))

	io.Copy(os.Stdout, resp.Body)
}
