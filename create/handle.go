package create

import (
	"fmt"
	"net/http"
)

//AccountCreator makes account
func AccountCreator(w http.ResponseWriter, r *http.Request) {
	msg := create(string([]byte(r.URL.Path)[6:]))
	fmt.Fprintf(w, "%v", msg)
	fmt.Printf("%v\n", msg)
}
