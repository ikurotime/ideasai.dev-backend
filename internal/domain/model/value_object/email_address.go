package value_object

type Email string

func NewEmailAddress(email string) Email {
	return Email(email)
}

func (e *Email) String() string {
	return string(*e)
}
