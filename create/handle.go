package create

import (
	"fmt"
	"net/http"
)

//FileRelatedMessage gives return message
type FileRelatedMessage struct {
	ReturnMessage string
}

var (
	fileAlreadyExists          = FileRelatedMessage{" file already exists "}
	errorInCreatingSentFile    = FileRelatedMessage{" error occurd while creating sent file "}
	errorInCreatingRecivedFile = FileRelatedMessage{" error occurd while creating recieved file "}
	doneMakingFolder           = FileRelatedMessage{" done !! "}
	errorInCreatingFriendFile  = FileRelatedMessage{" error occured while creating friends file "}
)

//AccountCreator makes account
func AccountCreator(w http.ResponseWriter, r *http.Request) {
	urlData := create(string([]byte(r.URL.Path)[6:]))
	fmt.Println(urlData.ReturnMessage + "\n")
	fmt.Printf("%v\n", r.Header)
	fmt.Fprintf(w, urlData.ReturnMessage)
}
