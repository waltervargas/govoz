package main

import (
	"flag"
	"log"

	"github.com/waltervargas/govoz"
)

func main() {
	client := flag.Bool("c", false, "starts a client")
	url := flag.String("url", "http://192.168.0.81:8080/audio", "URL to connect")
	flag.Parse()

	if *client {
		log.Printf("connecting to: %s\n", *url)
		err := govoz.RunAs("client", *url)
		if err != nil {
			panic(err)
		}
	}

	log.Println("running server")
	err := govoz.RunAs("server", "")
	if err != nil {
		panic(err)
	}
}
