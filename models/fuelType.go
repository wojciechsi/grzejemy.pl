package models

type FuelType struct {
	FuelName string
}

func (f FuelType) GetFuelType() string {
	return f.FuelName
}
