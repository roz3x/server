package posting

import (
	"fmt"
	"net/http"
)

//PostRelatedError gives return message
type PostRelatedError struct {
	ReturnMessage string
}

var (
	usersPath = "./users/"
)

var (
	nicelyDone               = PostRelatedError{" done !! sent yo message "}
	countaintthree           = PostRelatedError{" 3 arguements should be there "}
	senderaccountaintthere   = PostRelatedError{" cant find sender's account "}
	recieveraccountaintthere = PostRelatedError{" cant find reciever account "}
	notfriends               = PostRelatedError{" not friends "}
)

//Post 's the message
func Post(w http.ResponseWriter, r *http.Request) {
	t := post(string([]byte(r.URL.Path)[6:]))
	fmt.Printf(t.ReturnMessage + "\n")
	fmt.Fprintf(w, t.ReturnMessage)
}
