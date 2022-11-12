package main

import "fmt"

type FuelType int8

const (
	Groszek FuelType = iota //autonumber the rest
	Orzech
	Opony
)

type Userer interface {
	getName() string
}

type UserData struct { //struct to compose
	name, password string
}

func (u UserData) getName() string { //function on
	return u.name //composed struct
}

type Vendor struct { //composition #1
	UserData    //because no inher.
	salesPoints []string
}

type User struct { //composition #2
	UserData
	favouriteFuel FuelType
}

func main() {

	jacuśData := UserData{ //delare nested
		name:     "Jacek Szedel", //struct
		password: "1234",
	}

	jacuś := User{ //declare nesting
		UserData:      jacuśData, //class
		favouriteFuel: Opony,
	}

	var sklepyUJacka []string

	jacuśHandlarz := Vendor{
		UserData:    jacuśData,
		salesPoints: sklepyUJacka,
	}

	var users []Userer //polymorphic vector

	users = append(users, jacuś)
	users = append(users, jacuśHandlarz)

	fmt.Println(jacuś.getName()) //call method from
	//nested struct
	//as it would be
	//inherited

	fmt.Println(users[0].getName()) //getName must be declared
	//in interface definiction
}
