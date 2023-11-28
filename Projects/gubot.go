package main

import (
	"github.com/ArthurHlt/gubot/robot"

	// adapters
	_ "github.com/ArthurHlt/gubot/adapter/shell"

	// scripts
	_ "github.com/ArthurHlt/gubot/scripts"

	"log"
	"os"
)

func main() {
	addr := ":8080"
	port := os.Getenv("PORT")
	if port != "" {
		addr = ":" + port
	}
	log.Fatal(robot.Start(addr))
}
