package create

import (
	"os"
	"strings"
)

var (
	usersPath = "./users/"
)

func create(user string) int {
	fields := strings.Split(user, "/")
	if len(fields) != 2 {
		return -1
	}
	err := os.Mkdir(usersPath+fields[0], 0777)
	if err != nil {
		return -2
	}
	_, err = os.Create(usersPath + fields[0] + "/sent")
	if err != nil {
		return -3
	}
	f, err := os.Create(usersPath + fields[0] + "/passwd")
	if err != nil {
		return -4
	}
	_, err = f.Write([]byte(fields[1]))
	if err != nil {
		return -5
	}
	_, err = os.Create(usersPath + fields[0] + "/recieved")
	if err != nil {
		return -6
	}
	_, err = os.Create(usersPath + fields[0] + "/friends")
	if err != nil {
		return -7
	}
	return 0
}
