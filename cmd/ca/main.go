package main

import (
	"fmt"

	"github.com/sidhman/clean-architecture/internal/api/rest"
	"github.com/sidhman/clean-architecture/pkg/auth/memory"
	"github.com/sidhman/clean-architecture/pkg/logic/safe"
	"github.com/sidhman/clean-architecture/pkg/storage/fs"
)

func main() {
	p := map[string]memory.Permission{
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

	auth := memory.New(p)
	store := fs.New(".")
	c := safe.New(auth, store)

	hs := rest.New(c)
	fmt.Println(hs.Serve())
}
