// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: DeleteLockCommand)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockDeleteLockCommand struct {
	fail func(message string, callerSkip ...int)
}

func NewMockDeleteLockCommand(options ...pegomock.Option) *MockDeleteLockCommand {
	mock := &MockDeleteLockCommand{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockDeleteLockCommand) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockDeleteLockCommand) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockDeleteLockCommand) DeleteLock(_param0 string) (*models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockDeleteLockCommand().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DeleteLock", params, []reflect.Type{reflect.TypeOf((**models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockDeleteLockCommand) DeleteLocksByPull(_param0 string, _param1 int) (int, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockDeleteLockCommand().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DeleteLocksByPull", params, []reflect.Type{reflect.TypeOf((*int)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 int
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(int)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockDeleteLockCommand) VerifyWasCalledOnce() *VerifierMockDeleteLockCommand {
	return &VerifierMockDeleteLockCommand{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockDeleteLockCommand) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockDeleteLockCommand {
	return &VerifierMockDeleteLockCommand{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockDeleteLockCommand) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockDeleteLockCommand {
	return &VerifierMockDeleteLockCommand{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockDeleteLockCommand) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockDeleteLockCommand {
	return &VerifierMockDeleteLockCommand{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockDeleteLockCommand struct {
	mock                   *MockDeleteLockCommand
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockDeleteLockCommand) DeleteLock(_param0 string) *MockDeleteLockCommand_DeleteLock_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DeleteLock", params, verifier.timeout)
	return &MockDeleteLockCommand_DeleteLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockDeleteLockCommand_DeleteLock_OngoingVerification struct {
	mock              *MockDeleteLockCommand
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockDeleteLockCommand_DeleteLock_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockDeleteLockCommand_DeleteLock_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockDeleteLockCommand) DeleteLocksByPull(_param0 string, _param1 int) *MockDeleteLockCommand_DeleteLocksByPull_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DeleteLocksByPull", params, verifier.timeout)
	return &MockDeleteLockCommand_DeleteLocksByPull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockDeleteLockCommand_DeleteLocksByPull_OngoingVerification struct {
	mock              *MockDeleteLockCommand
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockDeleteLockCommand_DeleteLocksByPull_OngoingVerification) GetCapturedArguments() (string, int) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockDeleteLockCommand_DeleteLocksByPull_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
	}
	return
}
