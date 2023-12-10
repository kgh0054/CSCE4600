package builtins

import (
	"fmt"
	"os"
	"log"
)

func ListContents() error {
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}

	return err
}