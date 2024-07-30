package value_object

import "fmt"

type Name string

func NewName(name string) Name {
	return Name(name)
}

func (n *Name) String() string {
	return fmt.Sprintf("%s", *n)
}
