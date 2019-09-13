package posting

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	if check(internalformat[0], internalformat[1]) {
		fmt.Fprintf(senderFile, "%v--->%v\n", internalformat[1], internalformat[2])
		fmt.Fprintf(recieverFile, "%v--->%v\n", internalformat[0], internalformat[2])
	} else {
		return notfriends
	}
	return nicelyDone
}

func check(a, b string) bool {
	aFile, err := ioutil.ReadFile(usersPath + a + "/friends")
	if err != nil {
		return false
	}
	for _, slice := range strings.Split(string(aFile), "\n") {
		if slice == b {
			return true
		}
	}
	return false
}
