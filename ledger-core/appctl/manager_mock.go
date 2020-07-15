package appctl

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// ManagerMock implements Manager
type ManagerMock struct {
	t minimock.Tester

	funcCommitPulseChange          func(p1 PulseChange) (err error)
	inspectFuncCommitPulseChange   func(p1 PulseChange)
	afterCommitPulseChangeCounter  uint64
	beforeCommitPulseChangeCounter uint64
	CommitPulseChangeMock          mManagerMockCommitPulseChange
}

// NewManagerMock returns a mock for Manager
func NewManagerMock(t minimock.Tester) *ManagerMock {
	m := &ManagerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CommitPulseChangeMock = mManagerMockCommitPulseChange{mock: m}
	m.CommitPulseChangeMock.callArgs = []*ManagerMockCommitPulseChangeParams{}

	return m
}

type mManagerMockCommitPulseChange struct {
	mock               *ManagerMock
	defaultExpectation *ManagerMockCommitPulseChangeExpectation
	expectations       []*ManagerMockCommitPulseChangeExpectation

	callArgs []*ManagerMockCommitPulseChangeParams
	mutex    sync.RWMutex
}

// ManagerMockCommitPulseChangeExpectation specifies expectation struct of the Manager.CommitPulseChange
type ManagerMockCommitPulseChangeExpectation struct {
	mock    *ManagerMock
	params  *ManagerMockCommitPulseChangeParams
	results *ManagerMockCommitPulseChangeResults
	Counter uint64
}

// ManagerMockCommitPulseChangeParams contains parameters of the Manager.CommitPulseChange
type ManagerMockCommitPulseChangeParams struct {
	p1 PulseChange
}

// ManagerMockCommitPulseChangeResults contains results of the Manager.CommitPulseChange
type ManagerMockCommitPulseChangeResults struct {
	err error
}

// Expect sets up expected params for Manager.CommitPulseChange
func (mmCommitPulseChange *mManagerMockCommitPulseChange) Expect(p1 PulseChange) *mManagerMockCommitPulseChange {
	if mmCommitPulseChange.mock.funcCommitPulseChange != nil {
		mmCommitPulseChange.mock.t.Fatalf("ManagerMock.CommitPulseChange mock is already set by Set")
	}

	if mmCommitPulseChange.defaultExpectation == nil {
		mmCommitPulseChange.defaultExpectation = &ManagerMockCommitPulseChangeExpectation{}
	}

	mmCommitPulseChange.defaultExpectation.params = &ManagerMockCommitPulseChangeParams{p1}
	for _, e := range mmCommitPulseChange.expectations {
		if minimock.Equal(e.params, mmCommitPulseChange.defaultExpectation.params) {
			mmCommitPulseChange.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCommitPulseChange.defaultExpectation.params)
		}
	}

	return mmCommitPulseChange
}

// Inspect accepts an inspector function that has same arguments as the Manager.CommitPulseChange
func (mmCommitPulseChange *mManagerMockCommitPulseChange) Inspect(f func(p1 PulseChange)) *mManagerMockCommitPulseChange {
	if mmCommitPulseChange.mock.inspectFuncCommitPulseChange != nil {
		mmCommitPulseChange.mock.t.Fatalf("Inspect function is already set for ManagerMock.CommitPulseChange")
	}

	mmCommitPulseChange.mock.inspectFuncCommitPulseChange = f

	return mmCommitPulseChange
}

// Return sets up results that will be returned by Manager.CommitPulseChange
func (mmCommitPulseChange *mManagerMockCommitPulseChange) Return(err error) *ManagerMock {
	if mmCommitPulseChange.mock.funcCommitPulseChange != nil {
		mmCommitPulseChange.mock.t.Fatalf("ManagerMock.CommitPulseChange mock is already set by Set")
	}

	if mmCommitPulseChange.defaultExpectation == nil {
		mmCommitPulseChange.defaultExpectation = &ManagerMockCommitPulseChangeExpectation{mock: mmCommitPulseChange.mock}
	}
	mmCommitPulseChange.defaultExpectation.results = &ManagerMockCommitPulseChangeResults{err}
	return mmCommitPulseChange.mock
}

//Set uses given function f to mock the Manager.CommitPulseChange method
func (mmCommitPulseChange *mManagerMockCommitPulseChange) Set(f func(p1 PulseChange) (err error)) *ManagerMock {
	if mmCommitPulseChange.defaultExpectation != nil {
		mmCommitPulseChange.mock.t.Fatalf("Default expectation is already set for the Manager.CommitPulseChange method")
	}

	if len(mmCommitPulseChange.expectations) > 0 {
		mmCommitPulseChange.mock.t.Fatalf("Some expectations are already set for the Manager.CommitPulseChange method")
	}

	mmCommitPulseChange.mock.funcCommitPulseChange = f
	return mmCommitPulseChange.mock
}

// When sets expectation for the Manager.CommitPulseChange which will trigger the result defined by the following
// Then helper
func (mmCommitPulseChange *mManagerMockCommitPulseChange) When(p1 PulseChange) *ManagerMockCommitPulseChangeExpectation {
	if mmCommitPulseChange.mock.funcCommitPulseChange != nil {
		mmCommitPulseChange.mock.t.Fatalf("ManagerMock.CommitPulseChange mock is already set by Set")
	}

	expectation := &ManagerMockCommitPulseChangeExpectation{
		mock:   mmCommitPulseChange.mock,
		params: &ManagerMockCommitPulseChangeParams{p1},
	}
	mmCommitPulseChange.expectations = append(mmCommitPulseChange.expectations, expectation)
	return expectation
}

// Then sets up Manager.CommitPulseChange return parameters for the expectation previously defined by the When method
func (e *ManagerMockCommitPulseChangeExpectation) Then(err error) *ManagerMock {
	e.results = &ManagerMockCommitPulseChangeResults{err}
	return e.mock
}

// CommitPulseChange implements Manager
func (mmCommitPulseChange *ManagerMock) CommitPulseChange(p1 PulseChange) (err error) {
	mm_atomic.AddUint64(&mmCommitPulseChange.beforeCommitPulseChangeCounter, 1)
	defer mm_atomic.AddUint64(&mmCommitPulseChange.afterCommitPulseChangeCounter, 1)

	if mmCommitPulseChange.inspectFuncCommitPulseChange != nil {
		mmCommitPulseChange.inspectFuncCommitPulseChange(p1)
	}

	mm_params := &ManagerMockCommitPulseChangeParams{p1}

	// Record call args
	mmCommitPulseChange.CommitPulseChangeMock.mutex.Lock()
	mmCommitPulseChange.CommitPulseChangeMock.callArgs = append(mmCommitPulseChange.CommitPulseChangeMock.callArgs, mm_params)
	mmCommitPulseChange.CommitPulseChangeMock.mutex.Unlock()

	for _, e := range mmCommitPulseChange.CommitPulseChangeMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCommitPulseChange.CommitPulseChangeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCommitPulseChange.CommitPulseChangeMock.defaultExpectation.Counter, 1)
		mm_want := mmCommitPulseChange.CommitPulseChangeMock.defaultExpectation.params
		mm_got := ManagerMockCommitPulseChangeParams{p1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCommitPulseChange.t.Errorf("ManagerMock.CommitPulseChange got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCommitPulseChange.CommitPulseChangeMock.defaultExpectation.results
		if mm_results == nil {
			mmCommitPulseChange.t.Fatal("No results are set for the ManagerMock.CommitPulseChange")
		}
		return (*mm_results).err
	}
	if mmCommitPulseChange.funcCommitPulseChange != nil {
		return mmCommitPulseChange.funcCommitPulseChange(p1)
	}
	mmCommitPulseChange.t.Fatalf("Unexpected call to ManagerMock.CommitPulseChange. %v", p1)
	return
}

// CommitPulseChangeAfterCounter returns a count of finished ManagerMock.CommitPulseChange invocations
func (mmCommitPulseChange *ManagerMock) CommitPulseChangeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCommitPulseChange.afterCommitPulseChangeCounter)
}

// CommitPulseChangeBeforeCounter returns a count of ManagerMock.CommitPulseChange invocations
func (mmCommitPulseChange *ManagerMock) CommitPulseChangeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCommitPulseChange.beforeCommitPulseChangeCounter)
}

// Calls returns a list of arguments used in each call to ManagerMock.CommitPulseChange.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCommitPulseChange *mManagerMockCommitPulseChange) Calls() []*ManagerMockCommitPulseChangeParams {
	mmCommitPulseChange.mutex.RLock()

	argCopy := make([]*ManagerMockCommitPulseChangeParams, len(mmCommitPulseChange.callArgs))
	copy(argCopy, mmCommitPulseChange.callArgs)

	mmCommitPulseChange.mutex.RUnlock()

	return argCopy
}

// MinimockCommitPulseChangeDone returns true if the count of the CommitPulseChange invocations corresponds
// the number of defined expectations
func (m *ManagerMock) MinimockCommitPulseChangeDone() bool {
	for _, e := range m.CommitPulseChangeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CommitPulseChangeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCommitPulseChangeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCommitPulseChange != nil && mm_atomic.LoadUint64(&m.afterCommitPulseChangeCounter) < 1 {
		return false
	}
	return true
}

// MinimockCommitPulseChangeInspect logs each unmet expectation
func (m *ManagerMock) MinimockCommitPulseChangeInspect() {
	for _, e := range m.CommitPulseChangeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ManagerMock.CommitPulseChange with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CommitPulseChangeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCommitPulseChangeCounter) < 1 {
		if m.CommitPulseChangeMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ManagerMock.CommitPulseChange")
		} else {
			m.t.Errorf("Expected call to ManagerMock.CommitPulseChange with params: %#v", *m.CommitPulseChangeMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCommitPulseChange != nil && mm_atomic.LoadUint64(&m.afterCommitPulseChangeCounter) < 1 {
		m.t.Error("Expected call to ManagerMock.CommitPulseChange")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ManagerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCommitPulseChangeInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ManagerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ManagerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCommitPulseChangeDone()
}
