// Code generated by MockGen. DO NOT EDIT.
// Source: sorted_unique_list.go

// Package list is a generated GoMock package.
package list

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSortedUniqueList is a mock of SortedUniqueList interface.
type MockSortedUniqueList struct {
	ctrl     *gomock.Controller
	recorder *MockSortedUniqueListMockRecorder
}

// MockSortedUniqueListMockRecorder is the mock recorder for MockSortedUniqueList.
type MockSortedUniqueListMockRecorder struct {
	mock *MockSortedUniqueList
}

// NewMockSortedUniqueList creates a new mock instance.
func NewMockSortedUniqueList(ctrl *gomock.Controller) *MockSortedUniqueList {
	mock := &MockSortedUniqueList{ctrl: ctrl}
	mock.recorder = &MockSortedUniqueListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSortedUniqueList) EXPECT() *MockSortedUniqueListMockRecorder {
	return m.recorder
}

// Contains mocks base method.
func (m *MockSortedUniqueList) Contains(arg0 Item) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contains", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Contains indicates an expected call of Contains.
func (mr *MockSortedUniqueListMockRecorder) Contains(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contains", reflect.TypeOf((*MockSortedUniqueList)(nil).Contains), arg0)
}

// Insert mocks base method.
func (m *MockSortedUniqueList) Insert(arg0 Item) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Insert", arg0)
}

// Insert indicates an expected call of Insert.
func (mr *MockSortedUniqueListMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockSortedUniqueList)(nil).Insert), arg0)
}

// Len mocks base method.
func (m *MockSortedUniqueList) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len.
func (mr *MockSortedUniqueListMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockSortedUniqueList)(nil).Len))
}

// Range mocks base method.
func (m *MockSortedUniqueList) Range(arg0 func(Item) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Range", arg0)
}

// Range indicates an expected call of Range.
func (mr *MockSortedUniqueListMockRecorder) Range(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Range", reflect.TypeOf((*MockSortedUniqueList)(nil).Range), arg0)
}

// Remove mocks base method.
func (m *MockSortedUniqueList) Remove(arg0 Item) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove.
func (mr *MockSortedUniqueListMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockSortedUniqueList)(nil).Remove), arg0)
}

// ReverseRange mocks base method.
func (m *MockSortedUniqueList) ReverseRange(fn func(Item) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReverseRange", fn)
}

// ReverseRange indicates an expected call of ReverseRange.
func (mr *MockSortedUniqueListMockRecorder) ReverseRange(fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReverseRange", reflect.TypeOf((*MockSortedUniqueList)(nil).ReverseRange), fn)
}