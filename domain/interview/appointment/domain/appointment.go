package domain

import "github.com/mrbryside/rbh/domain/interview/appointment/types/mystatus"

type Appointment struct {
	Id          uint
	Name        string
	Description string
	Status      mystatus.Type
	Enabled     bool
	CreatedAt   string
}
