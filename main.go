package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/vcokltfre/projman/src/manager"
	"github.com/vcokltfre/projman/src/store"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Args must be provided.")
		os.Exit(1)
	}

	store := store.NewStore()
	manager := manager.ProjectManager{
		Store: *store,
	}

	manager.Init()

	command := args[0]
	switch command {
	case "start":
		if len(args) == 1 {
			wd, _ := os.Getwd()
			parts := strings.Split(wd, "/")
			name := parts[len(parts)-1]

			manager.Start(name)
		} else {
			manager.Start(args[1])
		}
	case "close":
		if len(args) == 1 {
			wd, _ := os.Getwd()

			manager.CloseFromLocation(wd)
		} else {
			manager.Close(args[1])
		}
	case "restore":
		manager.Unclose(args[1])
	case "list":
		manager.List()
	case "cleanup":
		manager.Cleanup(false)
	case "cleanup-all":
		manager.Cleanup(true)
	case "validate":
		manager.Validate()
	}
}
