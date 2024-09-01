// Code generated by mockery v2.43.0. DO NOT EDIT.

package mock

import (
	"github.com/joseasousa/desafio_cloud_run/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// ZipCodeService is an autogenerated mock type for the ZipCodeService type
type ZipCodeService struct {
	mock.Mock
}

type ZipCodeService_Expecter struct {
	mock *mock.Mock
}

func (_m *ZipCodeService) EXPECT() *ZipCodeService_Expecter {
	return &ZipCodeService_Expecter{mock: &_m.Mock}
}

// GetLocationByZipCode provides a mock function with given fields: zipCode
func (_m *ZipCodeService) GetLocationByZipCode(zipCode string) (*domain.Location, error) {
	ret := _m.Called(zipCode)

	if len(ret) == 0 {
		panic("no return value specified for GetLocationByZipCode")
	}

	var r0 *domain.Location
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Location, error)); ok {
		return rf(zipCode)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Location); ok {
		r0 = rf(zipCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Location)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(zipCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ZipCodeService_GetLocationByZipCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLocationByZipCode'
type ZipCodeService_GetLocationByZipCode_Call struct {
	*mock.Call
}

// GetLocationByZipCode is a helper method to define mock.On call
//   - zipCode string
func (_e *ZipCodeService_Expecter) GetLocationByZipCode(zipCode interface{}) *ZipCodeService_GetLocationByZipCode_Call {
	return &ZipCodeService_GetLocationByZipCode_Call{Call: _e.mock.On("GetLocationByZipCode", zipCode)}
}

func (_c *ZipCodeService_GetLocationByZipCode_Call) Run(run func(zipCode string)) *ZipCodeService_GetLocationByZipCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ZipCodeService_GetLocationByZipCode_Call) Return(_a0 *domain.Location, _a1 error) *ZipCodeService_GetLocationByZipCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ZipCodeService_GetLocationByZipCode_Call) RunAndReturn(run func(string) (*domain.Location, error)) *ZipCodeService_GetLocationByZipCode_Call {
	_c.Call.Return(run)
	return _c
}

// NewZipCodeService creates a new instance of ZipCodeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewZipCodeService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ZipCodeService {
	mock := &ZipCodeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
