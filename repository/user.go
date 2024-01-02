package repository

import (
	"github.com/3boku/Go-Server/types"
	"github.com/3boku/Go-Server/types/errors"
)

type UserRepository struct {
	userMap []*types.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{}, //user map초기화
	}
}

func (u *UserRepository) Create(user *types.User) error {
	u.userMap = append(u.userMap, user)
	return nil
}

func (u *UserRepository) Update(name string, newAge int64) error {
	isExisted := false

	for index, user := range u.userMap {
		if user.Name == name {
			user.Age = newAge
			u.userMap = append(u.userMap[:index], u.userMap[index+1:]...)

			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Delete(userName string) error {
	isExisted := false

	for index, user := range u.userMap {
		if user.Name == userName {
			u.userMap = append(u.userMap[:index], u.userMap[index+1:]...)

			isExisted = true
			continue
		}
	}

	if !isExisted {
		return errors.Errorf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}
