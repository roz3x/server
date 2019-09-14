package posting

import (
	"fmt"
	"net/http"
)

const (
	usersPath = "../users/"
)

//Post 's the message
func Post(w http.ResponseWriter, r *http.Request) {
	msg := post(string([]byte(r.URL.Path)[6:]))
	fmt.Printf("%v\n", msg)
	fmt.Fprintf(w, "%v", msg)
}
