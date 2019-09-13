package friend

import (
	"fmt"
	"net/http"
)

const (
	userPath = "./users/"
)

//Friendserror is the return type
type Friendserror struct {
	ReturnMessage string
}

var (
	nicelydone          = Friendserror{" nicely done adding 2 frinds "}
	countaintthree      = Friendserror{" there must be 2 arguements "}
	openingsendersfile  = Friendserror{" error while opening sender's friend file "}
	openingrecieverfile = Friendserror{" error while opening reciever's friend file "}
)

//AddFriend adds data to friend file
func AddFriend(w http.ResponseWriter, r *http.Request) {
	msg := add(string([]byte(r.URL.Path)[6:]))
	fmt.Fprintf(w, msg.ReturnMessage+"\n")
	fmt.Printf(msg.ReturnMessage + "\n")
}
