package friend

import (
	"fmt"
	"os"
	"strings"
)

func add(querry string) Friendserror {
	fields := strings.Split(querry, "/")
	if len(fields) != 2 {
		return countaintthree
	}
	from, err := os.OpenFile(userPath+fields[0]+"/friends", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return openingsendersfile
	}
	to, err := os.OpenFile(userPath+fields[1]+"/friends", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return openingrecieverfile
	}
	fmt.Fprintf(from, fields[1]+"\n")
	fmt.Fprintf(to, fields[0]+"\n")
	from.Close()
	to.Close()
	return nicelydone
}
