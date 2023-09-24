package main

import (
	"fmt"
	"os"

	"github.com/sidh/clean-architecture/internal/api/rest"
	"github.com/sidh/clean-architecture/internal/auth/memory"
	"github.com/sidh/clean-architecture/internal/logic"
	"github.com/sidh/clean-architecture/internal/logic/safe"
	"github.com/sidh/clean-architecture/internal/logic/unsafe"
	"github.com/sidh/clean-architecture/internal/storage"
	"github.com/sidh/clean-architecture/internal/storage/fs"
	storeMemory "github.com/sidh/clean-architecture/internal/storage/memory"
)

func main() {
	users := map[string]memory.Permission{
		"all": {
			Store: true,
			Load:  true,
		},
		"store": {
			Store: true,
			Load:  false,
		},
		"load": {
			Store: false,
			Load:  true,
		},
	}

	auth := memory.New(users)

	args := os.Args[1:]
	if len(args) < 2 {
		printHelp()
		return
	}

	var additionalArgs int

	var store storage.Storage
	switch args[0] {
	case "memory":
		store = storeMemory.New()
	case "fs":
		store = fs.New(args[1])
		additionalArgs++
	}

	if len(args) < 2+additionalArgs {
		printHelp()
		return
	}

	var core logic.Core
	switch args[1+additionalArgs] {
	case "safe":
		core = safe.New(auth, store)
	case "unsafe":
		core = unsafe.New(store)
	}

	hs := rest.New(core)
	fmt.Println(hs.Serve())
}

func printHelp() {
	fmt.Println("kv <storage> [storage_arguments] <logic>")
	fmt.Println("storage:")
	fmt.Println(" * memory")
	fmt.Println(" * fs <path_where_to_store_data>")
	fmt.Println("logic:")
	fmt.Println(" * safe")
	fmt.Println(" * unsafe")
}
