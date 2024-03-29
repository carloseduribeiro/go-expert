package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number int64 = 0

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		//number++
		//m.Unlock()
		atomic.AddInt64(&number, 1)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("You are the %d visitant\n", number)))
	})
	http.ListenAndServe(":8000", nil)
}
