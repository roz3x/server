package home

import (
	"fmt"
	"log"
	"net/http"
)

//Home is home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "hello")
	pusher, ok := w.(http.Pusher)
	if !ok {
		log.Fatal(ok)
		return
	}
	if err := pusher.Push("./README.md", nil); err != nil {
		log.Fatal(ok)
		return
	}
}
