package user

import (
	"github.com/Junx27/first-crud/entities"
	"github.com/Junx27/first-crud/repositories"
)

type useCase struct {
	userRepo repositories.UserRepository
}

func NewUseCase(userRepo repositories.UserRepository) useCase {
	return useCase{userRepo: userRepo}
}

type UserService interface {
	FindAll() ([]entities.User, error)
}

func (u useCase) FindAll() ([]entities.User, error) {
	return u.userRepo.FindAll()
}
