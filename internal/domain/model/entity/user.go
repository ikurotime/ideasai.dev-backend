package entity

import . "ikurotime/ideasai/internal/domain/model/value_object"

type User struct {
	ID    UserID
	Name  Name
	Email Email
}

func NewUser(ID UserID, Name Name, Email Email) *User {
	return &User{
		ID,
		Name,
		Email,
	}
}

func (u *User) String() string {
	return u.ID.String() + " " + u.Name.String()
}
func (u *User) GetID() UserID {
	return u.ID
}
func (u *User) GetName() Name {
	return u.Name
}
func (u *User) GetEmail() Email {
	return u.Email
}
func (u *User) SetName(name Name) {
	u.Name = name
}
func (u *User) SetEmail(email Email) {
	u.Email = email
}
