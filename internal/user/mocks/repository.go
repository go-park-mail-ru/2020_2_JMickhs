package mocks

import (
	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/user/models"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct{
	mock.Mock
}

func (_m *UserRepository) Add(user models.User) (models.User, error){
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


func (_m *UserRepository) GetUserByID(ID int) (models.User, error) {
	ret := _m.Called(ID)

	var r0 models.User
	if rf,ok := ret.Get(0).(func(ID int) models.User); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf,ok := ret.Get(1).(func(ID int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0,r1
}

func (_m *UserRepository) GetByUserName(name string) (models.User, error) {
	ret := _m.Called(name)

	var r0 models.User
	if rf,ok := ret.Get(0).(func(name string) models.User); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf,ok := ret.Get(1).(func(name string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0,r1
}

func (_m *UserRepository) UpdateAvatar(user models.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf,ok := ret.Get(0).(func(user models.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}


func (_m *UserRepository) UpdateUser(user models.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf,ok := ret.Get(0).(func(user models.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *UserRepository) UpdatePassword(user models.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf,ok := ret.Get(0).(func(user models.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}