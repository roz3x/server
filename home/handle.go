package home

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Home is home page
func Home(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile("./text.htm")
	if err != nil {
		log.Print("error with file part\n")
		return
	}
	fmt.Fprintf(w, string(body))
	push, ok := w.(http.Pusher)
	if !ok {
		return
	}
	push.Push("hello", nil)
}
