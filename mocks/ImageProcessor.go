// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ImageProcessor is an autogenerated mock type for the ImageProcessor type
type ImageProcessor struct {
	mock.Mock
}

// CompressImage provides a mock function with given fields: inputFileName, outputFileName
func (_m *ImageProcessor) CompressImage(inputFileName string, outputFileName string) error {
	ret := _m.Called(inputFileName, outputFileName)

	if len(ret) == 0 {
		panic("no return value specified for CompressImage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(inputFileName, outputFileName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConvertToJPEG provides a mock function with given fields: inputFileName, outputFileName
func (_m *ImageProcessor) ConvertToJPEG(inputFileName string, outputFileName string) error {
	ret := _m.Called(inputFileName, outputFileName)

	if len(ret) == 0 {
		panic("no return value specified for ConvertToJPEG")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(inputFileName, outputFileName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResizeImage provides a mock function with given fields: inputFileName, width, height, outputFileName
func (_m *ImageProcessor) ResizeImage(inputFileName string, width int, height int, outputFileName string) error {
	ret := _m.Called(inputFileName, width, height, outputFileName)

	if len(ret) == 0 {
		panic("no return value specified for ResizeImage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, int, string) error); ok {
		r0 = rf(inputFileName, width, height, outputFileName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewImageProcessor creates a new instance of ImageProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewImageProcessor(t interface {
	mock.TestingT
	Cleanup(func())
}) *ImageProcessor {
	mock := &ImageProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}