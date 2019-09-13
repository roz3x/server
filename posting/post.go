package posting

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
)

func post(format string) PostRelatedError {
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
	fmt.Fprintf(senderFile, "%v--->%v\n", internalformat[1], internalformat[2])
	fmt.Fprintf(recieverFile, "%v--->%v\n", internalformat[0], internalformat[2])
	senderFile.Close()
	recieverFile.Close()
	return nicelyDone
}

//Check checks the frind list
func Check(a, b string) bool {
	aFile, err := ioutil.ReadFile(usersPath + a + "/friends")
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(aFile))
	return true
}
