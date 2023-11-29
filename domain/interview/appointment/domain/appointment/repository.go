package appointment

//go:generate mockgen -source=repository.go -destination=../../pkg/generated/appointmentmock/repository.go -package=appointmentmock
type Repository interface {
	Create(Aggregate) (Aggregate, error)
	GetAll(page, pageSize uint) (Aggregates, error)
	GetById(appointmentId uint) (Aggregate, error)
	UpdateById(Aggregate) (Aggregate, error)
}
