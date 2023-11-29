package history

//go:generate mockgen -source=repository.go -destination=../../pkg/generated/historymock/repository.go -package=historymock
type Repository interface {
	Create(Aggregate) (Aggregate, error)
	GetAllByAppointmentId(appointmentId uint) ([]Aggregate, error)
}
