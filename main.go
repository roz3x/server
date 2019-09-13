package main

import (
	"localhostServer/create"
	"localhostServer/friend"
	"localhostServer/posting"
	"net/http"
	"os"
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
	http.HandleFunc("/frnd/", friend.AddFriend)
	http.ListenAndServe(":"+port, nil)
}
