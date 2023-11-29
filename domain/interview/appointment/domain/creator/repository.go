package creator

//go:generate mockgen -source=repository.go -destination=../../pkg/generated/creatormock/repository.go -package=creatormock
type Repository interface {
	GetById(id uint) (Aggregate, error)
}
