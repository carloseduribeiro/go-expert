package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Creating file
	f, err := os.Create("file.txt")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	// size, err := f.WriteString("Hello World!")
	size, err := f.Write([]byte("Writing data\nOther line"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("File creted! Size: %d bytes\n", size)

	// Reading file
	f1, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(f1))

	// Reading file with buffer
	fmt.Println("\nReading file...")
	f2, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// Deleting file
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
