package model

type User struct {
	id    int
	Name  string
	Email string
}

func GetUser(id int) (*User, error) {
	// todo
	return &User{}, nil
}

func GetUsers() []User {
	return []User{}
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) IsNew() bool {
	return u.id != 0
}

func (u *User) Save() error {
	// todo
	return nil
}

func (u *User) Delete() error {
	return nil
}
