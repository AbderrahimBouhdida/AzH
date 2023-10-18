package main

import (
	"fmt"

	"golang.org/x/sys/windows/svc"

	"https://github.com/AbderrahimBouhdida/AzH/cmd"
	"https://github.com/AbderrahimBouhdida/AzH/constants"
)

func main() {
	fmt.Printf("%s %s\n%s\n\n", constants.DisplayName, constants.Version, constants.AuthorRef)

	if isWinSvc, err := svc.IsWindowsService(); err != nil {
		panic(err)
	} else if isWinSvc {
		if err != nil {
			panic(err)
		}
	} else {
		cmd.Execute()
	}
}
