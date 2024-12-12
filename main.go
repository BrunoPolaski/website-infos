package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BrunoPolaski/website-infos/app"
)

func main() {
	fmt.Println("Starting to search!")

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
