package main

import (
	"comand_line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println()

	application := app.Generate()

	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
