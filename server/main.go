package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync/atomic"
	"time"
	"github.com/gosuri/uilive"
)
// RAM usage and current HTTP sessions count.
func hello(w http.ResponseWriter, req *http.Request) {
	ch <- true
	atomic.AddInt32(&count, 1)
	time.Sleep(time.Second * 120)
	_, _ = fmt.Fprintf(w, "hello\n")
	atomic.AddInt32(&count, -1)
	ch <- true

}

var count int32
var ch chan bool

func main() {
	ch = make(chan bool)
	go func() {
		var m runtime.MemStats
		var writer = uilive.New()
		writer.Start()
		for {
			<-ch
			_, _ = fmt.Fprintf(writer, "Current connections count: %d\n", atomic.LoadInt32(&count))
			runtime.ReadMemStats(&m)
			_, _ = fmt.Fprintf(writer, "Alloc = %v MiB\n", m.Alloc/1024/1024)
			_, _ = fmt.Fprintf(writer, "TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
			_, _ = fmt.Fprintf(writer, "Sys = %v MiB\n", m.Sys/1024/1024)
			_, _ = fmt.Fprintf(writer, "NumGC = %v\n", m.NumGC)

		}
	}()

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
