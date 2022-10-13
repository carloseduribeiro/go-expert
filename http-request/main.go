package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(res))

	f, err := os.Create("google.html")
	if err != nil {
		panic(err)
	}
	size, err := f.Write(res)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Writed %d bytes in file google.html", size)
}
