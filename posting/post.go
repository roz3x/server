package posting

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type postRelatedError struct {
	returnMessage string
}

var (
	usersPath = "./users/"
)

var (
	nicelyDone               = postRelatedError{" done !! sent yo message "}
	countaintthree           = postRelatedError{" 3 arguements should be there "}
	senderaccountaintthere   = postRelatedError{" cant find sender's account "}
	recieveraccountaintthere = postRelatedError{" cant find reciever account "}
)

func post(format string) postRelatedError {
	internalformat := strings.Split(format, "/")
	if len(internalformat) != 3 {
		return countaintthree
	}
	senderFile, err := os.OpenFile(usersPath+internalformat[0]+"/sent", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return senderaccountaintthere
	}
	recieverFile, err := os.OpenFile(usersPath+internalformat[1]+"/recieved", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return recieveraccountaintthere
	}
	fmt.Fprintf(senderFile, "%v--->%v--->%v\n", time.Now(), internalformat[1], internalformat[2])
	fmt.Fprintf(recieverFile, "%v--->%v--->%v\n", time.Now(), internalformat[0], internalformat[2])
	senderFile.Close()
	recieverFile.Close()
	return nicelyDone
}
