package main

import (
	"fmt"
	"CLI/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}