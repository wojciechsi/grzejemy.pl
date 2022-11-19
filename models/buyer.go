package models

type Buyer struct {
	UserData
	favouriteFuel FuelType
}

func NewBuyer(name string) User {
	return Buyer{
		UserData: newUserData(name),
	}
}

//Can I has cheezburger
