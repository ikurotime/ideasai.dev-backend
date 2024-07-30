package value_object

type Description struct {
	string
}

func NewDescription(description string) *Description {
	return &Description{
		description,
	}
}

func (d *Description) String() string {
	return d.string
}
