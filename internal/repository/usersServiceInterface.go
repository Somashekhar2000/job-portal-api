package repository

import (
	model "project/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source=usersServiceInterface.go -destination=usersServiceInterface_mock.go -package=repository
type UsersService interface {
	UserSignup(userSignup model.UserSignup) (model.User, error)
	Userlogin(userLogin model.UserLogin) (jwt.RegisteredClaims, error)
}
