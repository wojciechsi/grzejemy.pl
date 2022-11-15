package models

import "time"

type Offer struct {
	SalesPoint
	FuelType
	time.Time
	price float32
}
