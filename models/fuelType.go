package models

type FuelType struct {
	Name string
}

func (f FuelType) GetFuelType() string {
	return f.Name
}
