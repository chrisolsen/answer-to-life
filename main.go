package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":4200", nil)
}

func handleRoot(rw http.ResponseWriter, req *http.Request) {
	time.Sleep(5 * time.Second)
	rw.Write([]byte(`<div style="font-size:442px; text-align:center;">42</div>`))
}
