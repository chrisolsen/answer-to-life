package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"menteslibres.net/gosexy/redis"
)

var rClient *redis.Client

func main() {
	// redis setup
	rClient = redis.New()
	if err := rClient.Connect("redis", 6379); err != nil {
		log.Println("Unable to connect to redis")
		return
	}

	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":4200", nil)
}

func handleRoot(rw http.ResponseWriter, req *http.Request) {

	answer, err := rClient.Get("answer2life")
	if err == nil {
		rw.Write(renderAnswer(answer))
		return
	}

	time.Sleep(5 * time.Second)

	rClient.Set("answer2life", "42")
	rw.Write(renderAnswer("42"))
}

func renderAnswer(answer string) []byte {
	t := time.Now()
	template := `<div style="font-size:442px; text-align:center;">%v</div>
	<em>Calculated at %v</em>`

	return []byte(fmt.Sprintf(template, answer, t.String()))
}
