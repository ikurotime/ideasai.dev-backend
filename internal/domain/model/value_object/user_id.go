package valueobject

import "fmt"
 
type UserID int

func NewUserID(ID int) UserID {
     return UserID(ID)
}

func (u *UserID) String() string {
    return fmt.Sprintf("%d", *u)
}

func (u *UserID) Validate() bool {
    if *u == 0 {
        return false
    }
    return true
}
