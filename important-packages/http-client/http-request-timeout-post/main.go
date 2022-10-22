package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}
	jsonBody := bytes.NewBuffer([]byte(`{"name": "Carlos"}`))
	resp, err := c.Post(
		"https://www.google.com",
		"application/json",
		jsonBody,
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	_, err = io.CopyBuffer(os.Stdout, resp.Body, nil)
	if err != nil {
		panic(err)
	}
}
