package models

type Vendor struct {
	UserData    
	salesPoints []string
}

func NewVendor(name string) Vendor {
	return Vendor{
		UserData: newUserData(name),
	}
}