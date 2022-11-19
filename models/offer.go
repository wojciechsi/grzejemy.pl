package models

import (
	"fmt"
	"time"
)

type Offer struct {
	SalesPoint
	FuelType
	time.Time
	Price    float64
	Comments []Comment
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
	o.Comments = append(o.Comments, c)
}

func (o *Offer) GetPrice() string {
	return fmt.Sprintf("%8.2f", o.Price)
}

func (o *Offer) GetComments() []Comment {
	return o.Comments
}
