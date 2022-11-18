package main

import (
	"fmt"
	"github.com/wojciechsi/grzejemy.pl/models"
	"github.com/wojciechsi/grzejemy.pl/routers"
	"time"
)

func main() {
	go routers.RunServer() //in another thread

	jacuś := models.NewBuyer("Jacek")
	jacek := models.NewVendor("dr Szedel")
	jacex := models.NewSalesPoint("Gliwice", jacek)

	taniosprzedam := models.NewOffer(jacex, models.FuelType{Name: "groszek"}, 3560.75)

	komentarz := models.NewComment(jacuś, "nie polecam")
	taniosprzedam.AddComment(komentarz)

	fmt.Println(taniosprzedam.GetName() +
		" at " + taniosprzedam.GetRegion() +
		" sells " + taniosprzedam.GetFuelType() +
		" " + taniosprzedam.GetPrice() + " per tone" +
		" offered " + taniosprzedam.Time.String()) //tu jak widać nie ma sensu nazwyać wszystkich pól structa xD
	//fmt.Println(taniosprzedam.GetComments()[0].GetName() + " says: " +
	//	taniosprzedam.GetComments()[0].GetContent() + " and verification status is: " +
	//	strconv.FormatBool(taniosprzedam.GetComments()[0].IsVerified()))

	time.Sleep(30 * time.Second) //wait enought ime to visit localhost:8080
}
