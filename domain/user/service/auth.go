package service

import (
	"errors"

	"github.com/mrbryside/rbh/domain/user/domain/user"
	"github.com/mrbryside/rbh/domain/user/types/myrole"
)

//go:generate mockgen -source=./auth.go -destination=../pkg/generated/mockgen/auth_service.go -package=mockgen
type AuthServicer interface {
	Register(name, email, password string) (string, error)
	Login(email, password string) (string, error)
	GetById(id uint) (user.Aggregate, error)
}

type AuthService struct {
	userDomain user.Repository
	jwtService JwtServicer
}

func NewAuthService(ur user.Repository, jws JwtServicer) AuthService {
	return AuthService{
		userDomain: ur,
		jwtService: jws,
	}
}

func (a AuthService) Register(name string, email string, password string) (string, error) {
	userAgg, err := a.prepareRegister(name, email, password)
	if err != nil {
		return "", err
	}
	createdAgg, err := a.userDomain.Create(userAgg)
	if err != nil {
		return "", err
	}
	id := createdAgg.Person().Id
	role := createdAgg.Person().Role.Value

	return a.generateToken(id, role)
}

func (a AuthService) Login(email, password string) (string, error) {
	isAuthen, userAgg := a.userDomain.Authenticate(email, password)
	if !isAuthen {
		return "", errors.New("authentication failed (email or password is incorrect)")
	}
	role := userAgg.Person().Role.Value
	id := userAgg.Person().Id

	return a.generateToken(id, role)
}

func (a AuthService) GetById(id uint) (user.Aggregate, error) {
	return a.userDomain.GetById(id)
}

func (a AuthService) generateToken(id uint, role string) (string, error) {
	return a.jwtService.GenerateToken(id, role)
}

func (a AuthService) prepareRegister(name string, email string, password string) (user.Aggregate, error) {
	userAgg, err := user.New(name, email, password)
	if err != nil {
		return user.Aggregate{}, err
	}
	return userAgg.SetRole(myrole.PeopleTeam)
}
