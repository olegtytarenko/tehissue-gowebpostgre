package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)
	for {
		log.Println("Time init client")
		time.Sleep(time.Second * 2)
	}
}
