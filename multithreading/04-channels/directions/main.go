package main

// chan<- means this channel is send-only
func receive(name string, hello chan<- string) {
	hello <- name
}

// <-chan means this channel is receive-only
func read(data <-chan string) {
	println(<-data)
}

func main() {
	hello := make(chan string)
	go receive("Carlos", hello)
	read(hello)
}
