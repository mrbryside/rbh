package comment

//go:generate mockgen -source=repository.go -destination=../../pkg/generated/commentmock/repository.go -package=commentmock
type Repository interface {
	Create(Aggregate) (Aggregate, error)
	GetAllByAppointmentId(appointmentId uint) ([]Aggregate, error)
	GetById(id uint) (Aggregate, error)
	UpdateById(Aggregate) (Aggregate, error)
	DeleteById(id uint) error
}
