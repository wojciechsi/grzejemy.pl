package models

type FuelType int8

const (
	Groszek FuelType = iota 				//autonumber the rest
	Orzech
	Opony
)

type Userer interface {
	getName() string
}

type UserData struct { 						//struct to compose
	name, password string
}

func (u UserData) getName() string { 		//function on
	return u.name 							//composed struct
}

type Vendor struct { 						//composition #1
	UserData    							//because no inher.
	salesPoints []string
}

type User struct { 							//composition #2
	UserData
	favouriteFuel FuelType
}
