package create

import "os"

var (
	usersPath = "./users/"
)

func create(user string) FileRelatedMessage {
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
	_, err = os.Create(usersPath + user + "/friends")
	if err != nil {
		return errorInCreatingFriendFile
	}
	return doneMakingFolder
}
