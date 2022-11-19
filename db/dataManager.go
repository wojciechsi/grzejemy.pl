package db

import (
	"fmt"
	"strconv"

	"github.com/wojciechsi/grzejemy.pl/models"
)

func GetTestOffer() []models.Offer {
	//some data before DB will be created
	jacuś := models.NewBuyer("Jacek")
	jacek := models.NewVendor("dr Szedel")
	jacex := models.NewSalesPoint("Gliwice", jacek)

	taniosprzedam := models.NewOffer(jacex, models.FuelType{FuelName: "groszek"}, 3560.75)
	taniosprzedam2 := models.NewOffer(jacex, models.FuelType{FuelName: "drewno"}, 900.75)
	taniosprzedam3 := models.NewOffer(jacex, models.FuelType{FuelName: "ekogroszek"}, 5000.75)
	komentarz := models.NewComment(jacuś, "nie polecam")
	taniosprzedam.AddComment(komentarz)

	//print to console
	fmt.Println(taniosprzedam.GetName() +
		" at " + taniosprzedam.GetRegion() +
		" sells " + taniosprzedam.GetFuelType() +
		" " + taniosprzedam.GetPrice() + " per tone" +
		" offered " + taniosprzedam.Time.String()) //tu jak widać nie ma sensu nazwyać wszystkich pól structa xD
	fmt.Println(taniosprzedam.GetComments()[0].GetName() + " says: " +
		taniosprzedam.GetComments()[0].GetContent() + " and verification status is: " +
		strconv.FormatBool(taniosprzedam.GetComments()[0].IsVerified()))

	offers := []models.Offer{
		taniosprzedam,
		taniosprzedam2,
		taniosprzedam3,
	}
	return offers
}
