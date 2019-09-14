package friend

import (
	"fmt"
	"os"
	"strings"
)

func add(querry string) int {
	fields := strings.Split(querry, "/")
	if len(fields) != 2 {
		return -1
	}
	from, err := os.OpenFile(userPath+fields[0]+"/friends", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return -2
	}
	to, err := os.OpenFile(userPath+fields[1]+"/friends", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return -3
	}
	fmt.Fprintf(from, fields[1]+"\n")
	fmt.Fprintf(to, fields[0]+"\n")
	from.Close()
	to.Close()
	return 0
}
