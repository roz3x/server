package main

import (
	"net/http"
	"server/create"
	"server/friend"
	"server/home"
	"server/posting"
	"server/ws"
)

func main() {
	port := "8080"
	http.HandleFunc("/", home.Home)
	http.HandleFunc("/user/", create.AccountCreator)
	http.HandleFunc("/post/", posting.Post)
	http.HandleFunc("/frnd/", friend.AddFriend)
	http.HandleFunc("/ws", ws.Ws)
	http.ListenAndServe(":"+port, nil)
}
