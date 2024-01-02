package repository

import (
	"sync"
)

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	//repository 이곳에선 데이터베이스 같은것들을 설정해줌
	User *UserRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: NewUserRepository(),
		}
	})

	return repositoryInstance
}
