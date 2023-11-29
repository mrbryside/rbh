package user

//go:generate mockgen -source=./repository.go -destination=../../pkg/generated/mockgen/user_domain_repo.go -package=mockgen
type Repository interface {
	Create(Aggregate) (Aggregate, error)
	GetById(userId uint) (Aggregate, error)
	GetByEmail(email string) (Aggregate, error)
	Authenticate(email string, password string) (bool, Aggregate)
}
