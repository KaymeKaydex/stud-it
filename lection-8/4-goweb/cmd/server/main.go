package main

import (
	"context"
	"errors"
	"log"
	"os"

	"test/internal/app"
)

func main() {
	ctx := context.Background()

	application, err := app.New(ctx)
	if err != nil {
		log.Printf("cant create application core : %s ", err)

		os.Exit(2)
	}

	err = application.Run()
	if errors.Is(err, context.Canceled) {
		log.Println("successfully stopped application")

		os.Exit(0)
	}

	log.Printf("server stoped with err: %s", err)
	os.Exit(2)
}
