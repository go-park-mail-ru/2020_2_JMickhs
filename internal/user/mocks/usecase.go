package mocks

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/models"
	"github.com/stretchr/testify/mock"
)

type UserUsercase struct {
	mock.Mock
}

func (_m *UserUsercase) Add(user models.User) (models.User, error){
	ret := _m.Called(user)

	var r0 models.User
	if rf,ok := ret.Get(0).(func(models.User) models.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf,ok := ret.Get(1).(func(models.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0,r1
}

func (_m *UserUsercase) GetById(user models.User) (models.User, error){
	ret := _m.Called(user)

	var r0 models.User
	if rf,ok := ret.Get(0).(func(models.User) models.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf,ok := ret.Get(1).(func(models.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0,r1
}