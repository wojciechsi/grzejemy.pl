package main

import (
	"github.com/wojciechsi/grzejemy.pl/routers"
	"time"
)

func main() {
	go routers.RunServer() //in another thread

	time.Sleep(30 * time.Second) //wait enought ime to visit localhost:8080
}
