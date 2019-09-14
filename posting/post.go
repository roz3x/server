package posting

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func post(format string) PostRelatedError {
	fields := strings.Split(format, "/")
	if len(fields) != 3 {
		return countaintthree
	}
	senderFile, err := os.OpenFile(usersPath+fields[0]+"/sent", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return senderaccountaintthere
	}
	recieverFile, err := os.OpenFile(usersPath+fields[1]+"/recieved", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return recieveraccountaintthere
	}
	if check(fields[0], fields[1]) {
		fmt.Fprintf(senderFile, "%v--->%v\n", fields[1], fields[2])
		fmt.Fprintf(recieverFile, "%v--->%v\n", fields[0], fields[2])
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
