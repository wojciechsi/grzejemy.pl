package models

import (
	"fmt"
	"time"
)

type Offer struct {
	SalesPoint
	FuelType
	time.Time
	price    float64
	comments []Comment
}

func NewOffer(salesPoint SalesPoint, fuelType FuelType, price float64) Offer {
	return Offer{
		salesPoint,
		fuelType,
		time.Now(),
		price,
		make([]Comment, 0),
	}
}

func (o *Offer) AddComment(c Comment) {
	o.comments = append(o.comments, c)
}

func (o *Offer) GetPrice() string {
	return fmt.Sprintf("%8.2f", o.price)
}

func (o *Offer) GetComments() []Comment {
	return o.comments
}
