package service

import (
	"github.com/3boku/Go-Server/repository"
	"sync"
)

// Network와 Repository의 다리 역할

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	repository *repository.Repository
	//repository

	User *User
}

func NewService(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}

		serviceInstance.User = newUserService(rep.User)
	})

	return serviceInstance
}
