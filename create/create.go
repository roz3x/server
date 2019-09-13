package create

var (
	usersPath = "./users/"
)

type fileRelatedMessage struct {
	returnMessage string
}

var (
	fileAlreadyExists          = fileRelatedMessage{" file already exists "}
	errorInCreatingSentFile    = fileRelatedMessage{" error occurd while creating sent file "}
	errorInCreatingRecivedFile = fileRelatedMessage{" error occurd while creating recieved file "}
	doneMakingFolder           = fileRelatedMessage{" done !! "}
)
