package home

import (
	"net/http"
)

//Home is home page
//update this t web app that i made
func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./home/index.html")
}
