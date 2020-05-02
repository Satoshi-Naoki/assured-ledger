package jet

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/assured-ledger/ledger-core/v2/insolar"
	"github.com/insolar/assured-ledger/ledger-core/v2/reference"
)

// CoordinatorMock implements Coordinator
type CoordinatorMock struct {
	t minimock.Tester

	funcMe          func() (g1 reference.Global)
	inspectFuncMe   func()
	afterMeCounter  uint64
	beforeMeCounter uint64
	MeMock          mCoordinatorMockMe

	funcQueryRole          func(ctx context.Context, role insolar.DynamicRole, obj reference.Local, pulse insolar.PulseNumber) (ga1 []reference.Global, err error)
	inspectFuncQueryRole   func(ctx context.Context, role insolar.DynamicRole, obj reference.Local, pulse insolar.PulseNumber)
	afterQueryRoleCounter  uint64
	beforeQueryRoleCounter uint64
	QueryRoleMock          mCoordinatorMockQueryRole
}

// NewCoordinatorMock returns a mock for Coordinator
func NewCoordinatorMock(t minimock.Tester) *CoordinatorMock {
	m := &CoordinatorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.MeMock = mCoordinatorMockMe{mock: m}

	m.QueryRoleMock = mCoordinatorMockQueryRole{mock: m}
	m.QueryRoleMock.callArgs = []*CoordinatorMockQueryRoleParams{}

	return m
}

type mCoordinatorMockMe struct {
	mock               *CoordinatorMock
	defaultExpectation *CoordinatorMockMeExpectation
	expectations       []*CoordinatorMockMeExpectation
}

// CoordinatorMockMeExpectation specifies expectation struct of the Coordinator.Me
type CoordinatorMockMeExpectation struct {
	mock *CoordinatorMock

	results *CoordinatorMockMeResults
	Counter uint64
}

// CoordinatorMockMeResults contains results of the Coordinator.Me
type CoordinatorMockMeResults struct {
	g1 reference.Global
}

// Expect sets up expected params for Coordinator.Me
func (mmMe *mCoordinatorMockMe) Expect() *mCoordinatorMockMe {
	if mmMe.mock.funcMe != nil {
		mmMe.mock.t.Fatalf("CoordinatorMock.Me mock is already set by Set")
	}

	if mmMe.defaultExpectation == nil {
		mmMe.defaultExpectation = &CoordinatorMockMeExpectation{}
	}

	return mmMe
}

// Inspect accepts an inspector function that has same arguments as the Coordinator.Me
func (mmMe *mCoordinatorMockMe) Inspect(f func()) *mCoordinatorMockMe {
	if mmMe.mock.inspectFuncMe != nil {
		mmMe.mock.t.Fatalf("Inspect function is already set for CoordinatorMock.Me")
	}

	mmMe.mock.inspectFuncMe = f

	return mmMe
}

// Return sets up results that will be returned by Coordinator.Me
func (mmMe *mCoordinatorMockMe) Return(g1 reference.Global) *CoordinatorMock {
	if mmMe.mock.funcMe != nil {
		mmMe.mock.t.Fatalf("CoordinatorMock.Me mock is already set by Set")
	}

	if mmMe.defaultExpectation == nil {
		mmMe.defaultExpectation = &CoordinatorMockMeExpectation{mock: mmMe.mock}
	}
	mmMe.defaultExpectation.results = &CoordinatorMockMeResults{g1}
	return mmMe.mock
}

//Set uses given function f to mock the Coordinator.Me method
func (mmMe *mCoordinatorMockMe) Set(f func() (g1 reference.Global)) *CoordinatorMock {
	if mmMe.defaultExpectation != nil {
		mmMe.mock.t.Fatalf("Default expectation is already set for the Coordinator.Me method")
	}

	if len(mmMe.expectations) > 0 {
		mmMe.mock.t.Fatalf("Some expectations are already set for the Coordinator.Me method")
	}

	mmMe.mock.funcMe = f
	return mmMe.mock
}

// Me implements Coordinator
func (mmMe *CoordinatorMock) Me() (g1 reference.Global) {
	mm_atomic.AddUint64(&mmMe.beforeMeCounter, 1)
	defer mm_atomic.AddUint64(&mmMe.afterMeCounter, 1)

	if mmMe.inspectFuncMe != nil {
		mmMe.inspectFuncMe()
	}

	if mmMe.MeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmMe.MeMock.defaultExpectation.Counter, 1)

		mm_results := mmMe.MeMock.defaultExpectation.results
		if mm_results == nil {
			mmMe.t.Fatal("No results are set for the CoordinatorMock.Me")
		}
		return (*mm_results).g1
	}
	if mmMe.funcMe != nil {
		return mmMe.funcMe()
	}
	mmMe.t.Fatalf("Unexpected call to CoordinatorMock.Me.")
	return
}

// MeAfterCounter returns a count of finished CoordinatorMock.Me invocations
func (mmMe *CoordinatorMock) MeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMe.afterMeCounter)
}

// MeBeforeCounter returns a count of CoordinatorMock.Me invocations
func (mmMe *CoordinatorMock) MeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMe.beforeMeCounter)
}

// MinimockMeDone returns true if the count of the Me invocations corresponds
// the number of defined expectations
func (m *CoordinatorMock) MinimockMeDone() bool {
	for _, e := range m.MeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.MeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterMeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMe != nil && mm_atomic.LoadUint64(&m.afterMeCounter) < 1 {
		return false
	}
	return true
}

// MinimockMeInspect logs each unmet expectation
func (m *CoordinatorMock) MinimockMeInspect() {
	for _, e := range m.MeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to CoordinatorMock.Me")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.MeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterMeCounter) < 1 {
		m.t.Error("Expected call to CoordinatorMock.Me")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMe != nil && mm_atomic.LoadUint64(&m.afterMeCounter) < 1 {
		m.t.Error("Expected call to CoordinatorMock.Me")
	}
}

type mCoordinatorMockQueryRole struct {
	mock               *CoordinatorMock
	defaultExpectation *CoordinatorMockQueryRoleExpectation
	expectations       []*CoordinatorMockQueryRoleExpectation

	callArgs []*CoordinatorMockQueryRoleParams
	mutex    sync.RWMutex
}

// CoordinatorMockQueryRoleExpectation specifies expectation struct of the Coordinator.QueryRole
type CoordinatorMockQueryRoleExpectation struct {
	mock    *CoordinatorMock
	params  *CoordinatorMockQueryRoleParams
	results *CoordinatorMockQueryRoleResults
	Counter uint64
}

// CoordinatorMockQueryRoleParams contains parameters of the Coordinator.QueryRole
type CoordinatorMockQueryRoleParams struct {
	ctx   context.Context
	role  insolar.DynamicRole
	obj   reference.Local
	pulse insolar.PulseNumber
}

// CoordinatorMockQueryRoleResults contains results of the Coordinator.QueryRole
type CoordinatorMockQueryRoleResults struct {
	ga1 []reference.Global
	err error
}

// Expect sets up expected params for Coordinator.QueryRole
func (mmQueryRole *mCoordinatorMockQueryRole) Expect(ctx context.Context, role insolar.DynamicRole, obj reference.Local, pulse insolar.PulseNumber) *mCoordinatorMockQueryRole {
	if mmQueryRole.mock.funcQueryRole != nil {
		mmQueryRole.mock.t.Fatalf("CoordinatorMock.QueryRole mock is already set by Set")
	}

	if mmQueryRole.defaultExpectation == nil {
		mmQueryRole.defaultExpectation = &CoordinatorMockQueryRoleExpectation{}
	}

	mmQueryRole.defaultExpectation.params = &CoordinatorMockQueryRoleParams{ctx, role, obj, pulse}
	for _, e := range mmQueryRole.expectations {
		if minimock.Equal(e.params, mmQueryRole.defaultExpectation.params) {
			mmQueryRole.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmQueryRole.defaultExpectation.params)
		}
	}

	return mmQueryRole
}

// Inspect accepts an inspector function that has same arguments as the Coordinator.QueryRole
func (mmQueryRole *mCoordinatorMockQueryRole) Inspect(f func(ctx context.Context, role insolar.DynamicRole, obj reference.Local, pulse insolar.PulseNumber)) *mCoordinatorMockQueryRole {
	if mmQueryRole.mock.inspectFuncQueryRole != nil {
		mmQueryRole.mock.t.Fatalf("Inspect function is already set for CoordinatorMock.QueryRole")
	}

	mmQueryRole.mock.inspectFuncQueryRole = f

	return mmQueryRole
}

// Return sets up results that will be returned by Coordinator.QueryRole
func (mmQueryRole *mCoordinatorMockQueryRole) Return(ga1 []reference.Global, err error) *CoordinatorMock {
	if mmQueryRole.mock.funcQueryRole != nil {
		mmQueryRole.mock.t.Fatalf("CoordinatorMock.QueryRole mock is already set by Set")
	}

	if mmQueryRole.defaultExpectation == nil {
		mmQueryRole.defaultExpectation = &CoordinatorMockQueryRoleExpectation{mock: mmQueryRole.mock}
	}
	mmQueryRole.defaultExpectation.results = &CoordinatorMockQueryRoleResults{ga1, err}
	return mmQueryRole.mock
}

//Set uses given function f to mock the Coordinator.QueryRole method
func (mmQueryRole *mCoordinatorMockQueryRole) Set(f func(ctx context.Context, role insolar.DynamicRole, obj reference.Local, pulse insolar.PulseNumber) (ga1 []reference.Global, err error)) *CoordinatorMock {
	if mmQueryRole.defaultExpectation != nil {
		mmQueryRole.mock.t.Fatalf("Default expectation is already set for the Coordinator.QueryRole method")
	}

	if len(mmQueryRole.expectations) > 0 {
		mmQueryRole.mock.t.Fatalf("Some expectations are already set for the Coordinator.QueryRole method")
	}

	mmQueryRole.mock.funcQueryRole = f
	return mmQueryRole.mock
}

// When sets expectation for the Coordinator.QueryRole which will trigger the result defined by the following
// Then helper
func (mmQueryRole *mCoordinatorMockQueryRole) When(ctx context.Context, role insolar.DynamicRole, obj reference.Local, pulse insolar.PulseNumber) *CoordinatorMockQueryRoleExpectation {
	if mmQueryRole.mock.funcQueryRole != nil {
		mmQueryRole.mock.t.Fatalf("CoordinatorMock.QueryRole mock is already set by Set")
	}

	expectation := &CoordinatorMockQueryRoleExpectation{
		mock:   mmQueryRole.mock,
		params: &CoordinatorMockQueryRoleParams{ctx, role, obj, pulse},
	}
	mmQueryRole.expectations = append(mmQueryRole.expectations, expectation)
	return expectation
}

// Then sets up Coordinator.QueryRole return parameters for the expectation previously defined by the When method
func (e *CoordinatorMockQueryRoleExpectation) Then(ga1 []reference.Global, err error) *CoordinatorMock {
	e.results = &CoordinatorMockQueryRoleResults{ga1, err}
	return e.mock
}

// QueryRole implements Coordinator
func (mmQueryRole *CoordinatorMock) QueryRole(ctx context.Context, role insolar.DynamicRole, obj reference.Local, pulse insolar.PulseNumber) (ga1 []reference.Global, err error) {
	mm_atomic.AddUint64(&mmQueryRole.beforeQueryRoleCounter, 1)
	defer mm_atomic.AddUint64(&mmQueryRole.afterQueryRoleCounter, 1)

	if mmQueryRole.inspectFuncQueryRole != nil {
		mmQueryRole.inspectFuncQueryRole(ctx, role, obj, pulse)
	}

	mm_params := &CoordinatorMockQueryRoleParams{ctx, role, obj, pulse}

	// Record call args
	mmQueryRole.QueryRoleMock.mutex.Lock()
	mmQueryRole.QueryRoleMock.callArgs = append(mmQueryRole.QueryRoleMock.callArgs, mm_params)
	mmQueryRole.QueryRoleMock.mutex.Unlock()

	for _, e := range mmQueryRole.QueryRoleMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ga1, e.results.err
		}
	}

	if mmQueryRole.QueryRoleMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmQueryRole.QueryRoleMock.defaultExpectation.Counter, 1)
		mm_want := mmQueryRole.QueryRoleMock.defaultExpectation.params
		mm_got := CoordinatorMockQueryRoleParams{ctx, role, obj, pulse}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmQueryRole.t.Errorf("CoordinatorMock.QueryRole got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmQueryRole.QueryRoleMock.defaultExpectation.results
		if mm_results == nil {
			mmQueryRole.t.Fatal("No results are set for the CoordinatorMock.QueryRole")
		}
		return (*mm_results).ga1, (*mm_results).err
	}
	if mmQueryRole.funcQueryRole != nil {
		return mmQueryRole.funcQueryRole(ctx, role, obj, pulse)
	}
	mmQueryRole.t.Fatalf("Unexpected call to CoordinatorMock.QueryRole. %v %v %v %v", ctx, role, obj, pulse)
	return
}

// QueryRoleAfterCounter returns a count of finished CoordinatorMock.QueryRole invocations
func (mmQueryRole *CoordinatorMock) QueryRoleAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmQueryRole.afterQueryRoleCounter)
}

// QueryRoleBeforeCounter returns a count of CoordinatorMock.QueryRole invocations
func (mmQueryRole *CoordinatorMock) QueryRoleBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmQueryRole.beforeQueryRoleCounter)
}

// Calls returns a list of arguments used in each call to CoordinatorMock.QueryRole.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmQueryRole *mCoordinatorMockQueryRole) Calls() []*CoordinatorMockQueryRoleParams {
	mmQueryRole.mutex.RLock()

	argCopy := make([]*CoordinatorMockQueryRoleParams, len(mmQueryRole.callArgs))
	copy(argCopy, mmQueryRole.callArgs)

	mmQueryRole.mutex.RUnlock()

	return argCopy
}

// MinimockQueryRoleDone returns true if the count of the QueryRole invocations corresponds
// the number of defined expectations
func (m *CoordinatorMock) MinimockQueryRoleDone() bool {
	for _, e := range m.QueryRoleMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.QueryRoleMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterQueryRoleCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcQueryRole != nil && mm_atomic.LoadUint64(&m.afterQueryRoleCounter) < 1 {
		return false
	}
	return true
}

// MinimockQueryRoleInspect logs each unmet expectation
func (m *CoordinatorMock) MinimockQueryRoleInspect() {
	for _, e := range m.QueryRoleMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CoordinatorMock.QueryRole with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.QueryRoleMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterQueryRoleCounter) < 1 {
		if m.QueryRoleMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CoordinatorMock.QueryRole")
		} else {
			m.t.Errorf("Expected call to CoordinatorMock.QueryRole with params: %#v", *m.QueryRoleMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcQueryRole != nil && mm_atomic.LoadUint64(&m.afterQueryRoleCounter) < 1 {
		m.t.Error("Expected call to CoordinatorMock.QueryRole")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CoordinatorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockMeInspect()

		m.MinimockQueryRoleInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CoordinatorMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *CoordinatorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockMeDone() &&
		m.MinimockQueryRoleDone()
}
