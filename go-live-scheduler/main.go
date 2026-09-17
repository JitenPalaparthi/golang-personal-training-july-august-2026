package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type Event struct {
	Type string `json:"type"`
	G    int64  `json:"g"`
	P    int    `json:"p"`
	Info string `json:"info"`
}

var next int64
var mu sync.Mutex
var clients = map[chan Event]bool{}

func emit(e Event) {
	mu.Lock()
	defer mu.Unlock()
	for c := range clients {
		select {
		case c <- e:
		default:
		}
	}
}
func events(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	f := w.(http.Flusher)
	c := make(chan Event, 64)
	mu.Lock()
	clients[c] = true
	mu.Unlock()
	defer func() { mu.Lock(); delete(clients, c); mu.Unlock() }()
	for {
		select {
		case e := <-c:
			b, _ := json.Marshal(e)
			fmt.Fprintf(w, "data: %s\n\n", b)
			f.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
func worker(id int64, k string) {
	p := int(id % 2)
	emit(Event{"runnable", id, p, "created"})
	time.Sleep(time.Duration(150+rand.Intn(400)) * time.Millisecond)
	emit(Event{"running", id, p, "scheduled"})
	if k == "net" {
		time.Sleep(300 * time.Millisecond)
		emit(Event{"netwait", id, p, "real HTTP I/O: waiting / netpoller"})
		resp, e := http.Get("http://127.0.0.1:8080/ping")
		if e == nil {
			resp.Body.Close()
		}
		emit(Event{"runnable", id, p, "I/O ready"})
		time.Sleep(150 * time.Millisecond)
		emit(Event{"running", id, p, "rescheduled"})
	} else if k == "sleep" {
		emit(Event{"waiting", id, p, "timer wait"})
		time.Sleep(900 * time.Millisecond)
		emit(Event{"runnable", id, p, "timer fired"})
		time.Sleep(150 * time.Millisecond)
		emit(Event{"running", id, p, "rescheduled"})
	} else {
		end := time.Now().Add(800 * time.Millisecond)
		for time.Now().Before(end) {
		}
	}
	time.Sleep(250 * time.Millisecond)
	emit(Event{"done", id, p, "finished"})
}
func start(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	if n < 1 {
		n = 10
	}
	if n > 50 {
		n = 50
	}
	ks := []string{"net", "cpu", "sleep", "net", "cpu"}
	for i := 0; i < n; i++ {
		id := atomic.AddInt64(&next, 1)
		go worker(id, ks[i%len(ks)])
	}
	fmt.Fprint(w, "started")
}
func ping(w http.ResponseWriter, r *http.Request) {
	time.Sleep(900 * time.Millisecond)
	fmt.Fprint(w, "pong")
}
func main() {
	http.HandleFunc("/events", events)
	http.HandleFunc("/start", start)
	http.HandleFunc("/ping", ping)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	log.Println("Open http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
