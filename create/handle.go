package create

import (
	"fmt"
	"net/http"
	"os"
)

//AccountCreator makes account
func AccountCreator(w http.ResponseWriter, r *http.Request) {
	urlData := create(string([]byte(r.URL.Path)[13:]))
	fmt.Printf(urlData.returnMessage + "\n")
	fmt.Fprintf(w, urlData.returnMessage)

}
func create(user string) fileRelatedMessage {

	err := os.Mkdir(usersPath+user, 0777)
	if err != nil {
		return fileAlreadyExists
	}
	_, err = os.Create(usersPath + user + "/sent")
	if err != nil {
		return errorInCreatingSentFile
	}
	_, err = os.Create(usersPath + user + "/recieved")
	if err != nil {
		return errorInCreatingRecivedFile
	}
	return doneMakingFolder

}
