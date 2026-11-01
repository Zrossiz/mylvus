package main

import (
	"log"

	"github.com/Zrossiz/mylvus/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatalf("init app err: %v", err)

	}

	err = a.Start()
	if err != nil {
		log.Fatalf("start app err: %v", err)
	}

	err = a.Wait()
	if err != nil {
		log.Fatalf("stop app err: %v", err)
	}
}
