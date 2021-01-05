// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"github.com/goharbor/harbor/src/lib/q"
	policy "github.com/goharbor/harbor/src/pkg/retention/policy"
	mock "github.com/stretchr/testify/mock"

	retention "github.com/goharbor/harbor/src/pkg/retention"
)

// APIController is an autogenerated mock type for the APIController type
type APIController struct {
	mock.Mock
}

// CreateRetention provides a mock function with given fields: p
func (_m *APIController) CreateRetention(p *policy.Metadata) (int64, error) {
	ret := _m.Called(p)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*policy.Metadata) int64); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*policy.Metadata) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRetention provides a mock function with given fields: id
func (_m *APIController) DeleteRetention(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRetention provides a mock function with given fields: id
func (_m *APIController) GetRetention(id int64) (*policy.Metadata, error) {
	ret := _m.Called(id)

	var r0 *policy.Metadata
	if rf, ok := ret.Get(0).(func(int64) *policy.Metadata); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*policy.Metadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRetentionExec provides a mock function with given fields: eid
func (_m *APIController) GetRetentionExec(eid int64) (*retention.Execution, error) {
	ret := _m.Called(eid)

	var r0 *retention.Execution
	if rf, ok := ret.Get(0).(func(int64) *retention.Execution); ok {
		r0 = rf(eid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*retention.Execution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(eid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRetentionExecTaskLog provides a mock function with given fields: taskID
func (_m *APIController) GetRetentionExecTaskLog(taskID int64) ([]byte, error) {
	ret := _m.Called(taskID)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(int64) []byte); ok {
		r0 = rf(taskID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(taskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRetentionExecTask provides a mock function with given fields: taskID
func (_m *APIController) GetRetentionExecTask(taskID int64) (*retention.Task, error) {
	return &retention.Task{
		ID:          1,
		ExecutionID: 1,
	}, nil
}

// GetTotalOfRetentionExecTasks provides a mock function with given fields: executionID
func (_m *APIController) GetTotalOfRetentionExecTasks(executionID int64) (int64, error) {
	ret := _m.Called(executionID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(executionID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(executionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalOfRetentionExecs provides a mock function with given fields: policyID
func (_m *APIController) GetTotalOfRetentionExecs(policyID int64) (int64, error) {
	ret := _m.Called(policyID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(policyID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(policyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRetentionExecTasks provides a mock function with given fields: executionID, query
func (_m *APIController) ListRetentionExecTasks(executionID int64, query *q.Query) ([]*retention.Task, error) {
	ret := _m.Called(executionID, query)

	var r0 []*retention.Task
	if rf, ok := ret.Get(0).(func(int64, *q.Query) []*retention.Task); ok {
		r0 = rf(executionID, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*retention.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *q.Query) error); ok {
		r1 = rf(executionID, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRetentionExecs provides a mock function with given fields: policyID, query
func (_m *APIController) ListRetentionExecs(policyID int64, query *q.Query) ([]*retention.Execution, error) {
	ret := _m.Called(policyID, query)

	var r0 []*retention.Execution
	if rf, ok := ret.Get(0).(func(int64, *q.Query) []*retention.Execution); ok {
		r0 = rf(policyID, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*retention.Execution)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, *q.Query) error); ok {
		r1 = rf(policyID, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OperateRetentionExec provides a mock function with given fields: eid, action
func (_m *APIController) OperateRetentionExec(eid int64, action string) error {
	ret := _m.Called(eid, action)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, string) error); ok {
		r0 = rf(eid, action)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TriggerRetentionExec provides a mock function with given fields: policyID, trigger, dryRun
func (_m *APIController) TriggerRetentionExec(policyID int64, trigger string, dryRun bool) (int64, error) {
	ret := _m.Called(policyID, trigger, dryRun)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, string, bool) int64); ok {
		r0 = rf(policyID, trigger, dryRun)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string, bool) error); ok {
		r1 = rf(policyID, trigger, dryRun)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRetention provides a mock function with given fields: p
func (_m *APIController) UpdateRetention(p *policy.Metadata) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(*policy.Metadata) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
