package main

import(
	"github.com/wojciechsi/grzejemy.pl/controller"
	"github.com/wojciechsi/grzejemy.pl/models"
)

func main () {
	//controller.Run()
	jacuśData := UserData{ 					//delare nested
		name:     "Jacek Szedel", 			//struct
		password: "1234",
	}

	jacuś := User{ 							//declare nesting
		UserData:      jacuśData, 			//class
		favouriteFuel: Opony,
	}

	var sklepyUJacka []string

	jacuśHandlarz := Vendor{
		UserData:    jacuśData,
		salesPoints: sklepyUJacka,
	}

	var users []Userer 						//polymorphic vector

	users = append(users, jacuś)
	users = append(users, jacuśHandlarz)

	fmt.Println(jacuś.getName()) 			//call method from
											//nested struct
											//as it would be
											//inherited

	fmt.Println(users[0].getName()) 		//getName must be declared
											//in interface definiction
}