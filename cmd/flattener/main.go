package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"os/signal"
	"tangelo-flattener/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	// start the server.
	go serv.Start()

	// Wait for an in interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown.
	err = serv.Close()
	if err != nil {
		log.Fatal(err)
	}
}
