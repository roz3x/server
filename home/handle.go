package home

import (
	"fmt"
	"net/http"
)

//Home is home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "hello")
	pusher, ok := w.(http.Pusher)
	if !ok {
		return
	}
	if err := pusher.Push("./README.md", nil); err != nil {
		return
	}
}
