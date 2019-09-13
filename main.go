package main

import (
	"localhostServer/posting"
	"net/http"
	"os"

	"localhostServer/create"
)

var (
	f        *os.File
	err      error
	bodyFile string
)

func main() {
	port := "8080"
	http.HandleFunc("/user/create/", create.AccountCreator)
	http.HandleFunc("/post/", posting.Post)
	http.ListenAndServe(":"+port, nil)
}
