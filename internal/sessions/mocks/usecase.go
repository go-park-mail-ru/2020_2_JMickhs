package SessionMocks

import (
	"github.com/stretchr/testify/mock"
)

type SessionsUsecase struct {
	mock.Mock
}

func (_m *SessionsUsecase) AddToken(ID int) (string, error) {
	ret := _m.Called(ID)

	var r0 string
	if rf, ok := ret.Get(0).(func(ID int) string); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ID int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *SessionsUsecase) GetIDByToken(token string) (int, error) {
	ret := _m.Called(token)

	var r0 int
	if rf, ok := ret.Get(0).(func(token string) int); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(token string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *SessionsUsecase) DeleteSession(token string) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(token string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
