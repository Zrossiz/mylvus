package main

import (
	"log"
	"os"

	"github.com/Zrossiz/mylvus/internal/app"
)

func main() {
	err := app.Start()
	if err != nil {
		log.Printf("start app err: %v", err)
		os.Exit(1)
	}
}
