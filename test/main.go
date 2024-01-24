package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/White-Whale-Defi-Platform/migaloo-chain/v4/client/docs/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	err = fs.Walk(statikFS, "/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("statik path: ", path) // This will log the path of each file in the statikFS
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
