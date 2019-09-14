package friend

import (
	"fmt"
	"net/http"
)

const (
	userPath = "./users/"
)

//AddFriend adds data to friend file
func AddFriend(w http.ResponseWriter, r *http.Request) {
	msg := add(string([]byte(r.URL.Path)[6:]))
	fmt.Fprintf(w, "%v", msg)
	fmt.Printf("%v\n", msg)
}
