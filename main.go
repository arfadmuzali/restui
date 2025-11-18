package main

import (
	"log"

	"github.com/arfadmuzali/restui/cmd/restui"
)

func main() {
	err := restui.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
