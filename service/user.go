package service

import (
	"github.com/3boku/Go-Server/repository"
	"github.com/3boku/Go-Server/types"
)

type User struct {
	userRepository *repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (u *User) Create(newUser *types.User) error {
	return u.userRepository.Create(newUser)
}

func (u *User) Update(name string, newAge int64) error {
	return u.userRepository.Update(name, newAge)
}

func (u *User) Delete(user *types.User) error {
	return u.userRepository.Delete(user.Name)
}

func (u *User) Get() []*types.User {
	return u.userRepository.Get()
}
