package network

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	mm_beat "github.com/insolar/assured-ledger/ledger-core/appctl/beat"
	"github.com/insolar/assured-ledger/ledger-core/pulse"
)

// AccessorMock implements beat.Accessor
type AccessorMock struct {
	t minimock.Tester

	funcLatest          func(ctx context.Context) (b1 mm_beat.Beat, err error)
	inspectFuncLatest   func(ctx context.Context)
	afterLatestCounter  uint64
	beforeLatestCounter uint64
	LatestMock          mAccessorMockLatest

	funcOf          func(ctx context.Context, n1 pulse.Number) (b1 mm_beat.Beat, err error)
	inspectFuncOf   func(ctx context.Context, n1 pulse.Number)
	afterOfCounter  uint64
	beforeOfCounter uint64
	OfMock          mAccessorMockOf
}

// NewAccessorMock returns a mock for beat.Accessor
func NewAccessorMock(t minimock.Tester) *AccessorMock {
	m := &AccessorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.LatestMock = mAccessorMockLatest{mock: m}
	m.LatestMock.callArgs = []*AccessorMockLatestParams{}

	m.OfMock = mAccessorMockOf{mock: m}
	m.OfMock.callArgs = []*AccessorMockOfParams{}

	return m
}

type mAccessorMockLatest struct {
	mock               *AccessorMock
	defaultExpectation *AccessorMockLatestExpectation
	expectations       []*AccessorMockLatestExpectation

	callArgs []*AccessorMockLatestParams
	mutex    sync.RWMutex
}

// AccessorMockLatestExpectation specifies expectation struct of the Accessor.Latest
type AccessorMockLatestExpectation struct {
	mock    *AccessorMock
	params  *AccessorMockLatestParams
	results *AccessorMockLatestResults
	Counter uint64
}

// AccessorMockLatestParams contains parameters of the Accessor.Latest
type AccessorMockLatestParams struct {
	ctx context.Context
}

// AccessorMockLatestResults contains results of the Accessor.Latest
type AccessorMockLatestResults struct {
	b1  mm_beat.Beat
	err error
}

// Expect sets up expected params for Accessor.Latest
func (mmLatest *mAccessorMockLatest) Expect(ctx context.Context) *mAccessorMockLatest {
	if mmLatest.mock.funcLatest != nil {
		mmLatest.mock.t.Fatalf("AccessorMock.Latest mock is already set by Set")
	}

	if mmLatest.defaultExpectation == nil {
		mmLatest.defaultExpectation = &AccessorMockLatestExpectation{}
	}

	mmLatest.defaultExpectation.params = &AccessorMockLatestParams{ctx}
	for _, e := range mmLatest.expectations {
		if minimock.Equal(e.params, mmLatest.defaultExpectation.params) {
			mmLatest.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmLatest.defaultExpectation.params)
		}
	}

	return mmLatest
}

// Inspect accepts an inspector function that has same arguments as the Accessor.Latest
func (mmLatest *mAccessorMockLatest) Inspect(f func(ctx context.Context)) *mAccessorMockLatest {
	if mmLatest.mock.inspectFuncLatest != nil {
		mmLatest.mock.t.Fatalf("Inspect function is already set for AccessorMock.Latest")
	}

	mmLatest.mock.inspectFuncLatest = f

	return mmLatest
}

// Return sets up results that will be returned by Accessor.Latest
func (mmLatest *mAccessorMockLatest) Return(b1 mm_beat.Beat, err error) *AccessorMock {
	if mmLatest.mock.funcLatest != nil {
		mmLatest.mock.t.Fatalf("AccessorMock.Latest mock is already set by Set")
	}

	if mmLatest.defaultExpectation == nil {
		mmLatest.defaultExpectation = &AccessorMockLatestExpectation{mock: mmLatest.mock}
	}
	mmLatest.defaultExpectation.results = &AccessorMockLatestResults{b1, err}
	return mmLatest.mock
}

//Set uses given function f to mock the Accessor.Latest method
func (mmLatest *mAccessorMockLatest) Set(f func(ctx context.Context) (b1 mm_beat.Beat, err error)) *AccessorMock {
	if mmLatest.defaultExpectation != nil {
		mmLatest.mock.t.Fatalf("Default expectation is already set for the Accessor.Latest method")
	}

	if len(mmLatest.expectations) > 0 {
		mmLatest.mock.t.Fatalf("Some expectations are already set for the Accessor.Latest method")
	}

	mmLatest.mock.funcLatest = f
	return mmLatest.mock
}

// When sets expectation for the Accessor.Latest which will trigger the result defined by the following
// Then helper
func (mmLatest *mAccessorMockLatest) When(ctx context.Context) *AccessorMockLatestExpectation {
	if mmLatest.mock.funcLatest != nil {
		mmLatest.mock.t.Fatalf("AccessorMock.Latest mock is already set by Set")
	}

	expectation := &AccessorMockLatestExpectation{
		mock:   mmLatest.mock,
		params: &AccessorMockLatestParams{ctx},
	}
	mmLatest.expectations = append(mmLatest.expectations, expectation)
	return expectation
}

// Then sets up Accessor.Latest return parameters for the expectation previously defined by the When method
func (e *AccessorMockLatestExpectation) Then(b1 mm_beat.Beat, err error) *AccessorMock {
	e.results = &AccessorMockLatestResults{b1, err}
	return e.mock
}

// Latest implements beat.Accessor
func (mmLatest *AccessorMock) Latest(ctx context.Context) (b1 mm_beat.Beat, err error) {
	mm_atomic.AddUint64(&mmLatest.beforeLatestCounter, 1)
	defer mm_atomic.AddUint64(&mmLatest.afterLatestCounter, 1)

	if mmLatest.inspectFuncLatest != nil {
		mmLatest.inspectFuncLatest(ctx)
	}

	mm_params := &AccessorMockLatestParams{ctx}

	// Record call args
	mmLatest.LatestMock.mutex.Lock()
	mmLatest.LatestMock.callArgs = append(mmLatest.LatestMock.callArgs, mm_params)
	mmLatest.LatestMock.mutex.Unlock()

	for _, e := range mmLatest.LatestMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmLatest.LatestMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmLatest.LatestMock.defaultExpectation.Counter, 1)
		mm_want := mmLatest.LatestMock.defaultExpectation.params
		mm_got := AccessorMockLatestParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmLatest.t.Errorf("AccessorMock.Latest got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmLatest.LatestMock.defaultExpectation.results
		if mm_results == nil {
			mmLatest.t.Fatal("No results are set for the AccessorMock.Latest")
		}
		return (*mm_results).b1, (*mm_results).err
	}
	if mmLatest.funcLatest != nil {
		return mmLatest.funcLatest(ctx)
	}
	mmLatest.t.Fatalf("Unexpected call to AccessorMock.Latest. %v", ctx)
	return
}

// LatestAfterCounter returns a count of finished AccessorMock.Latest invocations
func (mmLatest *AccessorMock) LatestAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLatest.afterLatestCounter)
}

// LatestBeforeCounter returns a count of AccessorMock.Latest invocations
func (mmLatest *AccessorMock) LatestBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmLatest.beforeLatestCounter)
}

// Calls returns a list of arguments used in each call to AccessorMock.Latest.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmLatest *mAccessorMockLatest) Calls() []*AccessorMockLatestParams {
	mmLatest.mutex.RLock()

	argCopy := make([]*AccessorMockLatestParams, len(mmLatest.callArgs))
	copy(argCopy, mmLatest.callArgs)

	mmLatest.mutex.RUnlock()

	return argCopy
}

// MinimockLatestDone returns true if the count of the Latest invocations corresponds
// the number of defined expectations
func (m *AccessorMock) MinimockLatestDone() bool {
	for _, e := range m.LatestMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LatestMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLatestCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLatest != nil && mm_atomic.LoadUint64(&m.afterLatestCounter) < 1 {
		return false
	}
	return true
}

// MinimockLatestInspect logs each unmet expectation
func (m *AccessorMock) MinimockLatestInspect() {
	for _, e := range m.LatestMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AccessorMock.Latest with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.LatestMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterLatestCounter) < 1 {
		if m.LatestMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to AccessorMock.Latest")
		} else {
			m.t.Errorf("Expected call to AccessorMock.Latest with params: %#v", *m.LatestMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcLatest != nil && mm_atomic.LoadUint64(&m.afterLatestCounter) < 1 {
		m.t.Error("Expected call to AccessorMock.Latest")
	}
}

type mAccessorMockOf struct {
	mock               *AccessorMock
	defaultExpectation *AccessorMockOfExpectation
	expectations       []*AccessorMockOfExpectation

	callArgs []*AccessorMockOfParams
	mutex    sync.RWMutex
}

// AccessorMockOfExpectation specifies expectation struct of the Accessor.Of
type AccessorMockOfExpectation struct {
	mock    *AccessorMock
	params  *AccessorMockOfParams
	results *AccessorMockOfResults
	Counter uint64
}

// AccessorMockOfParams contains parameters of the Accessor.Of
type AccessorMockOfParams struct {
	ctx context.Context
	n1  pulse.Number
}

// AccessorMockOfResults contains results of the Accessor.Of
type AccessorMockOfResults struct {
	b1  mm_beat.Beat
	err error
}

// Expect sets up expected params for Accessor.Of
func (mmOf *mAccessorMockOf) Expect(ctx context.Context, n1 pulse.Number) *mAccessorMockOf {
	if mmOf.mock.funcOf != nil {
		mmOf.mock.t.Fatalf("AccessorMock.Of mock is already set by Set")
	}

	if mmOf.defaultExpectation == nil {
		mmOf.defaultExpectation = &AccessorMockOfExpectation{}
	}

	mmOf.defaultExpectation.params = &AccessorMockOfParams{ctx, n1}
	for _, e := range mmOf.expectations {
		if minimock.Equal(e.params, mmOf.defaultExpectation.params) {
			mmOf.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmOf.defaultExpectation.params)
		}
	}

	return mmOf
}

// Inspect accepts an inspector function that has same arguments as the Accessor.Of
func (mmOf *mAccessorMockOf) Inspect(f func(ctx context.Context, n1 pulse.Number)) *mAccessorMockOf {
	if mmOf.mock.inspectFuncOf != nil {
		mmOf.mock.t.Fatalf("Inspect function is already set for AccessorMock.Of")
	}

	mmOf.mock.inspectFuncOf = f

	return mmOf
}

// Return sets up results that will be returned by Accessor.Of
func (mmOf *mAccessorMockOf) Return(b1 mm_beat.Beat, err error) *AccessorMock {
	if mmOf.mock.funcOf != nil {
		mmOf.mock.t.Fatalf("AccessorMock.Of mock is already set by Set")
	}

	if mmOf.defaultExpectation == nil {
		mmOf.defaultExpectation = &AccessorMockOfExpectation{mock: mmOf.mock}
	}
	mmOf.defaultExpectation.results = &AccessorMockOfResults{b1, err}
	return mmOf.mock
}

//Set uses given function f to mock the Accessor.Of method
func (mmOf *mAccessorMockOf) Set(f func(ctx context.Context, n1 pulse.Number) (b1 mm_beat.Beat, err error)) *AccessorMock {
	if mmOf.defaultExpectation != nil {
		mmOf.mock.t.Fatalf("Default expectation is already set for the Accessor.Of method")
	}

	if len(mmOf.expectations) > 0 {
		mmOf.mock.t.Fatalf("Some expectations are already set for the Accessor.Of method")
	}

	mmOf.mock.funcOf = f
	return mmOf.mock
}

// When sets expectation for the Accessor.Of which will trigger the result defined by the following
// Then helper
func (mmOf *mAccessorMockOf) When(ctx context.Context, n1 pulse.Number) *AccessorMockOfExpectation {
	if mmOf.mock.funcOf != nil {
		mmOf.mock.t.Fatalf("AccessorMock.Of mock is already set by Set")
	}

	expectation := &AccessorMockOfExpectation{
		mock:   mmOf.mock,
		params: &AccessorMockOfParams{ctx, n1},
	}
	mmOf.expectations = append(mmOf.expectations, expectation)
	return expectation
}

// Then sets up Accessor.Of return parameters for the expectation previously defined by the When method
func (e *AccessorMockOfExpectation) Then(b1 mm_beat.Beat, err error) *AccessorMock {
	e.results = &AccessorMockOfResults{b1, err}
	return e.mock
}

// Of implements beat.Accessor
func (mmOf *AccessorMock) Of(ctx context.Context, n1 pulse.Number) (b1 mm_beat.Beat, err error) {
	mm_atomic.AddUint64(&mmOf.beforeOfCounter, 1)
	defer mm_atomic.AddUint64(&mmOf.afterOfCounter, 1)

	if mmOf.inspectFuncOf != nil {
		mmOf.inspectFuncOf(ctx, n1)
	}

	mm_params := &AccessorMockOfParams{ctx, n1}

	// Record call args
	mmOf.OfMock.mutex.Lock()
	mmOf.OfMock.callArgs = append(mmOf.OfMock.callArgs, mm_params)
	mmOf.OfMock.mutex.Unlock()

	for _, e := range mmOf.OfMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmOf.OfMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmOf.OfMock.defaultExpectation.Counter, 1)
		mm_want := mmOf.OfMock.defaultExpectation.params
		mm_got := AccessorMockOfParams{ctx, n1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmOf.t.Errorf("AccessorMock.Of got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmOf.OfMock.defaultExpectation.results
		if mm_results == nil {
			mmOf.t.Fatal("No results are set for the AccessorMock.Of")
		}
		return (*mm_results).b1, (*mm_results).err
	}
	if mmOf.funcOf != nil {
		return mmOf.funcOf(ctx, n1)
	}
	mmOf.t.Fatalf("Unexpected call to AccessorMock.Of. %v %v", ctx, n1)
	return
}

// OfAfterCounter returns a count of finished AccessorMock.Of invocations
func (mmOf *AccessorMock) OfAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmOf.afterOfCounter)
}

// OfBeforeCounter returns a count of AccessorMock.Of invocations
func (mmOf *AccessorMock) OfBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmOf.beforeOfCounter)
}

// Calls returns a list of arguments used in each call to AccessorMock.Of.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmOf *mAccessorMockOf) Calls() []*AccessorMockOfParams {
	mmOf.mutex.RLock()

	argCopy := make([]*AccessorMockOfParams, len(mmOf.callArgs))
	copy(argCopy, mmOf.callArgs)

	mmOf.mutex.RUnlock()

	return argCopy
}

// MinimockOfDone returns true if the count of the Of invocations corresponds
// the number of defined expectations
func (m *AccessorMock) MinimockOfDone() bool {
	for _, e := range m.OfMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.OfMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterOfCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcOf != nil && mm_atomic.LoadUint64(&m.afterOfCounter) < 1 {
		return false
	}
	return true
}

// MinimockOfInspect logs each unmet expectation
func (m *AccessorMock) MinimockOfInspect() {
	for _, e := range m.OfMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AccessorMock.Of with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.OfMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterOfCounter) < 1 {
		if m.OfMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to AccessorMock.Of")
		} else {
			m.t.Errorf("Expected call to AccessorMock.Of with params: %#v", *m.OfMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcOf != nil && mm_atomic.LoadUint64(&m.afterOfCounter) < 1 {
		m.t.Error("Expected call to AccessorMock.Of")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *AccessorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockLatestInspect()

		m.MinimockOfInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *AccessorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *AccessorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockLatestDone() &&
		m.MinimockOfDone()
}
