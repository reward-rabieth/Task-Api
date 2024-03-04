package main

import (
	"github.con/reward-rabieth/Task-Api/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
