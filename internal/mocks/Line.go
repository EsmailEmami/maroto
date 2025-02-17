// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	internal "github.com/johnfercher/maroto/internal"
	mock "github.com/stretchr/testify/mock"

	props "github.com/johnfercher/maroto/pkg/props"
)

// Line is an autogenerated mock type for the Line type
type Line struct {
	mock.Mock
}

// Draw provides a mock function with given fields: cell, lineProp
func (_m *Line) Draw(cell internal.Cell, lineProp props.Line) {
	_m.Called(cell, lineProp)
}

// NewLine creates a new instance of Line. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLine(t interface {
	mock.TestingT
	Cleanup(func())
}) *Line {
	mock := &Line{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
