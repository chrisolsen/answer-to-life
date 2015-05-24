package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":4200", nil)
}

func handleRoot(rw http.ResponseWriter, req *http.Request) {
	t := time.Now()
	time.Sleep(5 * time.Second)

	template := `<div style="font-size:442px; text-align:center;">42</div>
	<em>Calculated at %v</em>`

	content := []byte(fmt.Sprintf(template, t.String()))

	rw.Write(content)
}
