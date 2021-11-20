package main

import (
	"github.com/EQEmu/eqemu-docs-v2/internal/console"
	"github.com/EQEmu/eqemu-docs-v2/internal/http"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// if args passed, run console
	if len(os.Args) > 1 {
		err := console.Run()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	// if no args passed, run web server
	http.Run()
}
