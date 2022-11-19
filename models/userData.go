package models

type UserData struct {
	Name, password string
}

func newUserData(name string) UserData {
	return UserData{
		Name: name,
	}
}

func (u UserData) GetName() string {
	return u.Name
}
