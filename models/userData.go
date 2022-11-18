package models

type UserData struct {
	name, password string
}

func newUserData(name string) UserData {
	return UserData{
		name: name,
	}
}

func (u UserData) GetName() string {
	return u.name
}
