package value_object

import "fmt"

type Category string

func NewCategory(category string) Category {
	return Category(category)
}

func (c *Category) String() string {
	return fmt.Sprintf("%s", *c)
}
