package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0
	go func() {
		for {
			time.Sleep(time.Second * 2)
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 1)
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()
	//for i := 0; i < 3; i++ {
	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("Received from RabbitMQ: ID: %d - %s\n", msg1.id, msg1.msg)
		case msg2 := <-c2:
			fmt.Printf("Received from Kafka: ID: %d - %s\n", msg2.id, msg2.msg)
		case <-time.After(time.Second * 3):
			println("timout")
			//default:
			//	println("default")
		}
	}
}
