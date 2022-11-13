package main

import(
	"github.com/wojciechsi/grzejemy.pl/models"
	"github.com/wojciechsi/grzejemy.pl/controllers"
	"fmt"
	"time"
)

func main () {
	go controllers.RunServer() //in another thread

	jacuś := models.NewBuyer("Jacek")
	jacek := models.NewVendor("dr Szedel")

	var users []models.User

	users = append(users, jacek)
	users = append(users, jacuś)

	for i := 0; i<len(users); i++ {
		fmt.Println(users[i].GetName())
	}

	time.Sleep(30 * time.Second) //wait enought time to visit localhost:8080
}