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
		return -100
	}
	err := os.Mkdir(usersPath+user, 0777)
	if err != nil {
		return -2
	}
	_, err = os.Create(usersPath + user + "/sent")
	if err != nil {
		return -3
	}
	_, err = os.Create(usersPath + user + "/passwd")
	if err != nil {
		return -4
	}
	_, err = os.Create(usersPath + user + "/recieved")
	if err != nil {
		return -5
	}
	_, err = os.Create(usersPath + user + "/friends")
	if err != nil {
		return -6
	}
	return 0
}
