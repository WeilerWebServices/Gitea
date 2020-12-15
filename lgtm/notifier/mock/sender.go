package mocks

import mock "github.com/stretchr/testify/mock"
import notifier "github.com/go-gitea/lgtm/notifier"

// Sender is an autogenerated mock type for the Sender type
type Sender struct {
	mock.Mock
}

// Send provides a mock function with given fields: _a0
func (_m *Sender) Send(_a0 *notifier.Notification) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*notifier.Notification) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ notifier.Sender = (*Sender)(nil)