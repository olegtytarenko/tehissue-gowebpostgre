package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	server   *string
	port     *int
	endpoint *string

	pathToConfig *string
)

func init() {
	port = flag.Int("port", 80, "HTTP Port")
	server = flag.String("server", "localhost", "HTTP Host name or IP Address")
	endpoint = flag.String("endpoint", "api", "Endpoint `/api`, you can set param in config file")
	pathToConfig = flag.String("config", "", "Path to configuration file in format `json`")
}

func main() {
	// todo
	// web-service

	fmt.Println("Hello World", server, port, endpoint, pathToConfig)

	log.SetOutput(os.Stdout)
	for {
		log.Println("Time init")
		time.Sleep(time.Second * 2)
	}
}
