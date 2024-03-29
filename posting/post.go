package posting

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func post(format string) int {
	fields := strings.Split(format, "/")
	if len(fields) != 3 {
		return -1
	}
	senderFile, err := os.OpenFile(usersPath+fields[0]+"/sent", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return -2
	}
	recieverFile, err := os.OpenFile(usersPath+fields[1]+"/recieved", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return -3
	}
	if check(fields[0], fields[1]) {
		fmt.Fprintf(senderFile, "%v--->%v\n", fields[1], fields[2])
		fmt.Fprintf(recieverFile, "%v--->%v\n", fields[0], fields[2])
	} else {
		return -4
	}
	return 0
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
