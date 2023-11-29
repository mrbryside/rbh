package domain

import (
	"github.com/mrbryside/rbh/domain/user/types/myemail"
	"github.com/mrbryside/rbh/domain/user/types/mypass"
	"github.com/mrbryside/rbh/domain/user/types/myrole"
)

// Entity
type Person struct {
	Id       uint
	Name     string
	Email    myemail.Type
	Password mypass.Type
	Role     myrole.Type
}
