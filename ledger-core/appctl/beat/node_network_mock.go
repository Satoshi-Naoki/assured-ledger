package beat

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/insolar/assured-ledger/ledger-core/network/consensus/gcpv2/api/member"
	"github.com/insolar/assured-ledger/ledger-core/pulse"
	"github.com/insolar/assured-ledger/ledger-core/reference"
)

// NodeNetworkMock implements NodeNetwork
type NodeNetworkMock struct {
	t minimock.Tester

	funcGetAnyLatestNodeSnapshot          func() (n1 NodeSnapshot)
	inspectFuncGetAnyLatestNodeSnapshot   func()
	afterGetAnyLatestNodeSnapshotCounter  uint64
	beforeGetAnyLatestNodeSnapshotCounter uint64
	GetAnyLatestNodeSnapshotMock          mNodeNetworkMockGetAnyLatestNodeSnapshot

	funcGetLocalNodeReference          func() (h1 reference.Holder)
	inspectFuncGetLocalNodeReference   func()
	afterGetLocalNodeReferenceCounter  uint64
	beforeGetLocalNodeReferenceCounter uint64
	GetLocalNodeReferenceMock          mNodeNetworkMockGetLocalNodeReference

	funcGetLocalNodeRole          func() (p1 member.PrimaryRole)
	inspectFuncGetLocalNodeRole   func()
	afterGetLocalNodeRoleCounter  uint64
	beforeGetLocalNodeRoleCounter uint64
	GetLocalNodeRoleMock          mNodeNetworkMockGetLocalNodeRole

	funcGetNodeSnapshot          func(n1 pulse.Number) (n2 NodeSnapshot)
	inspectFuncGetNodeSnapshot   func(n1 pulse.Number)
	afterGetNodeSnapshotCounter  uint64
	beforeGetNodeSnapshotCounter uint64
	GetNodeSnapshotMock          mNodeNetworkMockGetNodeSnapshot
}

// NewNodeNetworkMock returns a mock for NodeNetwork
func NewNodeNetworkMock(t minimock.Tester) *NodeNetworkMock {
	m := &NodeNetworkMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetAnyLatestNodeSnapshotMock = mNodeNetworkMockGetAnyLatestNodeSnapshot{mock: m}

	m.GetLocalNodeReferenceMock = mNodeNetworkMockGetLocalNodeReference{mock: m}

	m.GetLocalNodeRoleMock = mNodeNetworkMockGetLocalNodeRole{mock: m}

	m.GetNodeSnapshotMock = mNodeNetworkMockGetNodeSnapshot{mock: m}
	m.GetNodeSnapshotMock.callArgs = []*NodeNetworkMockGetNodeSnapshotParams{}

	return m
}

type mNodeNetworkMockGetAnyLatestNodeSnapshot struct {
	mock               *NodeNetworkMock
	defaultExpectation *NodeNetworkMockGetAnyLatestNodeSnapshotExpectation
	expectations       []*NodeNetworkMockGetAnyLatestNodeSnapshotExpectation
}

// NodeNetworkMockGetAnyLatestNodeSnapshotExpectation specifies expectation struct of the NodeNetwork.GetAnyLatestNodeSnapshot
type NodeNetworkMockGetAnyLatestNodeSnapshotExpectation struct {
	mock *NodeNetworkMock

	results *NodeNetworkMockGetAnyLatestNodeSnapshotResults
	Counter uint64
}

// NodeNetworkMockGetAnyLatestNodeSnapshotResults contains results of the NodeNetwork.GetAnyLatestNodeSnapshot
type NodeNetworkMockGetAnyLatestNodeSnapshotResults struct {
	n1 NodeSnapshot
}

// Expect sets up expected params for NodeNetwork.GetAnyLatestNodeSnapshot
func (mmGetAnyLatestNodeSnapshot *mNodeNetworkMockGetAnyLatestNodeSnapshot) Expect() *mNodeNetworkMockGetAnyLatestNodeSnapshot {
	if mmGetAnyLatestNodeSnapshot.mock.funcGetAnyLatestNodeSnapshot != nil {
		mmGetAnyLatestNodeSnapshot.mock.t.Fatalf("NodeNetworkMock.GetAnyLatestNodeSnapshot mock is already set by Set")
	}

	if mmGetAnyLatestNodeSnapshot.defaultExpectation == nil {
		mmGetAnyLatestNodeSnapshot.defaultExpectation = &NodeNetworkMockGetAnyLatestNodeSnapshotExpectation{}
	}

	return mmGetAnyLatestNodeSnapshot
}

// Inspect accepts an inspector function that has same arguments as the NodeNetwork.GetAnyLatestNodeSnapshot
func (mmGetAnyLatestNodeSnapshot *mNodeNetworkMockGetAnyLatestNodeSnapshot) Inspect(f func()) *mNodeNetworkMockGetAnyLatestNodeSnapshot {
	if mmGetAnyLatestNodeSnapshot.mock.inspectFuncGetAnyLatestNodeSnapshot != nil {
		mmGetAnyLatestNodeSnapshot.mock.t.Fatalf("Inspect function is already set for NodeNetworkMock.GetAnyLatestNodeSnapshot")
	}

	mmGetAnyLatestNodeSnapshot.mock.inspectFuncGetAnyLatestNodeSnapshot = f

	return mmGetAnyLatestNodeSnapshot
}

// Return sets up results that will be returned by NodeNetwork.GetAnyLatestNodeSnapshot
func (mmGetAnyLatestNodeSnapshot *mNodeNetworkMockGetAnyLatestNodeSnapshot) Return(n1 NodeSnapshot) *NodeNetworkMock {
	if mmGetAnyLatestNodeSnapshot.mock.funcGetAnyLatestNodeSnapshot != nil {
		mmGetAnyLatestNodeSnapshot.mock.t.Fatalf("NodeNetworkMock.GetAnyLatestNodeSnapshot mock is already set by Set")
	}

	if mmGetAnyLatestNodeSnapshot.defaultExpectation == nil {
		mmGetAnyLatestNodeSnapshot.defaultExpectation = &NodeNetworkMockGetAnyLatestNodeSnapshotExpectation{mock: mmGetAnyLatestNodeSnapshot.mock}
	}
	mmGetAnyLatestNodeSnapshot.defaultExpectation.results = &NodeNetworkMockGetAnyLatestNodeSnapshotResults{n1}
	return mmGetAnyLatestNodeSnapshot.mock
}

//Set uses given function f to mock the NodeNetwork.GetAnyLatestNodeSnapshot method
func (mmGetAnyLatestNodeSnapshot *mNodeNetworkMockGetAnyLatestNodeSnapshot) Set(f func() (n1 NodeSnapshot)) *NodeNetworkMock {
	if mmGetAnyLatestNodeSnapshot.defaultExpectation != nil {
		mmGetAnyLatestNodeSnapshot.mock.t.Fatalf("Default expectation is already set for the NodeNetwork.GetAnyLatestNodeSnapshot method")
	}

	if len(mmGetAnyLatestNodeSnapshot.expectations) > 0 {
		mmGetAnyLatestNodeSnapshot.mock.t.Fatalf("Some expectations are already set for the NodeNetwork.GetAnyLatestNodeSnapshot method")
	}

	mmGetAnyLatestNodeSnapshot.mock.funcGetAnyLatestNodeSnapshot = f
	return mmGetAnyLatestNodeSnapshot.mock
}

// GetAnyLatestNodeSnapshot implements NodeNetwork
func (mmGetAnyLatestNodeSnapshot *NodeNetworkMock) GetAnyLatestNodeSnapshot() (n1 NodeSnapshot) {
	mm_atomic.AddUint64(&mmGetAnyLatestNodeSnapshot.beforeGetAnyLatestNodeSnapshotCounter, 1)
	defer mm_atomic.AddUint64(&mmGetAnyLatestNodeSnapshot.afterGetAnyLatestNodeSnapshotCounter, 1)

	if mmGetAnyLatestNodeSnapshot.inspectFuncGetAnyLatestNodeSnapshot != nil {
		mmGetAnyLatestNodeSnapshot.inspectFuncGetAnyLatestNodeSnapshot()
	}

	if mmGetAnyLatestNodeSnapshot.GetAnyLatestNodeSnapshotMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetAnyLatestNodeSnapshot.GetAnyLatestNodeSnapshotMock.defaultExpectation.Counter, 1)

		mm_results := mmGetAnyLatestNodeSnapshot.GetAnyLatestNodeSnapshotMock.defaultExpectation.results
		if mm_results == nil {
			mmGetAnyLatestNodeSnapshot.t.Fatal("No results are set for the NodeNetworkMock.GetAnyLatestNodeSnapshot")
		}
		return (*mm_results).n1
	}
	if mmGetAnyLatestNodeSnapshot.funcGetAnyLatestNodeSnapshot != nil {
		return mmGetAnyLatestNodeSnapshot.funcGetAnyLatestNodeSnapshot()
	}
	mmGetAnyLatestNodeSnapshot.t.Fatalf("Unexpected call to NodeNetworkMock.GetAnyLatestNodeSnapshot.")
	return
}

// GetAnyLatestNodeSnapshotAfterCounter returns a count of finished NodeNetworkMock.GetAnyLatestNodeSnapshot invocations
func (mmGetAnyLatestNodeSnapshot *NodeNetworkMock) GetAnyLatestNodeSnapshotAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetAnyLatestNodeSnapshot.afterGetAnyLatestNodeSnapshotCounter)
}

// GetAnyLatestNodeSnapshotBeforeCounter returns a count of NodeNetworkMock.GetAnyLatestNodeSnapshot invocations
func (mmGetAnyLatestNodeSnapshot *NodeNetworkMock) GetAnyLatestNodeSnapshotBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetAnyLatestNodeSnapshot.beforeGetAnyLatestNodeSnapshotCounter)
}

// MinimockGetAnyLatestNodeSnapshotDone returns true if the count of the GetAnyLatestNodeSnapshot invocations corresponds
// the number of defined expectations
func (m *NodeNetworkMock) MinimockGetAnyLatestNodeSnapshotDone() bool {
	for _, e := range m.GetAnyLatestNodeSnapshotMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetAnyLatestNodeSnapshotMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetAnyLatestNodeSnapshotCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetAnyLatestNodeSnapshot != nil && mm_atomic.LoadUint64(&m.afterGetAnyLatestNodeSnapshotCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetAnyLatestNodeSnapshotInspect logs each unmet expectation
func (m *NodeNetworkMock) MinimockGetAnyLatestNodeSnapshotInspect() {
	for _, e := range m.GetAnyLatestNodeSnapshotMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to NodeNetworkMock.GetAnyLatestNodeSnapshot")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetAnyLatestNodeSnapshotMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetAnyLatestNodeSnapshotCounter) < 1 {
		m.t.Error("Expected call to NodeNetworkMock.GetAnyLatestNodeSnapshot")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetAnyLatestNodeSnapshot != nil && mm_atomic.LoadUint64(&m.afterGetAnyLatestNodeSnapshotCounter) < 1 {
		m.t.Error("Expected call to NodeNetworkMock.GetAnyLatestNodeSnapshot")
	}
}

type mNodeNetworkMockGetLocalNodeReference struct {
	mock               *NodeNetworkMock
	defaultExpectation *NodeNetworkMockGetLocalNodeReferenceExpectation
	expectations       []*NodeNetworkMockGetLocalNodeReferenceExpectation
}

// NodeNetworkMockGetLocalNodeReferenceExpectation specifies expectation struct of the NodeNetwork.GetLocalNodeReference
type NodeNetworkMockGetLocalNodeReferenceExpectation struct {
	mock *NodeNetworkMock

	results *NodeNetworkMockGetLocalNodeReferenceResults
	Counter uint64
}

// NodeNetworkMockGetLocalNodeReferenceResults contains results of the NodeNetwork.GetLocalNodeReference
type NodeNetworkMockGetLocalNodeReferenceResults struct {
	h1 reference.Holder
}

// Expect sets up expected params for NodeNetwork.GetLocalNodeReference
func (mmGetLocalNodeReference *mNodeNetworkMockGetLocalNodeReference) Expect() *mNodeNetworkMockGetLocalNodeReference {
	if mmGetLocalNodeReference.mock.funcGetLocalNodeReference != nil {
		mmGetLocalNodeReference.mock.t.Fatalf("NodeNetworkMock.GetLocalNodeReference mock is already set by Set")
	}

	if mmGetLocalNodeReference.defaultExpectation == nil {
		mmGetLocalNodeReference.defaultExpectation = &NodeNetworkMockGetLocalNodeReferenceExpectation{}
	}

	return mmGetLocalNodeReference
}

// Inspect accepts an inspector function that has same arguments as the NodeNetwork.GetLocalNodeReference
func (mmGetLocalNodeReference *mNodeNetworkMockGetLocalNodeReference) Inspect(f func()) *mNodeNetworkMockGetLocalNodeReference {
	if mmGetLocalNodeReference.mock.inspectFuncGetLocalNodeReference != nil {
		mmGetLocalNodeReference.mock.t.Fatalf("Inspect function is already set for NodeNetworkMock.GetLocalNodeReference")
	}

	mmGetLocalNodeReference.mock.inspectFuncGetLocalNodeReference = f

	return mmGetLocalNodeReference
}

// Return sets up results that will be returned by NodeNetwork.GetLocalNodeReference
func (mmGetLocalNodeReference *mNodeNetworkMockGetLocalNodeReference) Return(h1 reference.Holder) *NodeNetworkMock {
	if mmGetLocalNodeReference.mock.funcGetLocalNodeReference != nil {
		mmGetLocalNodeReference.mock.t.Fatalf("NodeNetworkMock.GetLocalNodeReference mock is already set by Set")
	}

	if mmGetLocalNodeReference.defaultExpectation == nil {
		mmGetLocalNodeReference.defaultExpectation = &NodeNetworkMockGetLocalNodeReferenceExpectation{mock: mmGetLocalNodeReference.mock}
	}
	mmGetLocalNodeReference.defaultExpectation.results = &NodeNetworkMockGetLocalNodeReferenceResults{h1}
	return mmGetLocalNodeReference.mock
}

//Set uses given function f to mock the NodeNetwork.GetLocalNodeReference method
func (mmGetLocalNodeReference *mNodeNetworkMockGetLocalNodeReference) Set(f func() (h1 reference.Holder)) *NodeNetworkMock {
	if mmGetLocalNodeReference.defaultExpectation != nil {
		mmGetLocalNodeReference.mock.t.Fatalf("Default expectation is already set for the NodeNetwork.GetLocalNodeReference method")
	}

	if len(mmGetLocalNodeReference.expectations) > 0 {
		mmGetLocalNodeReference.mock.t.Fatalf("Some expectations are already set for the NodeNetwork.GetLocalNodeReference method")
	}

	mmGetLocalNodeReference.mock.funcGetLocalNodeReference = f
	return mmGetLocalNodeReference.mock
}

// GetLocalNodeReference implements NodeNetwork
func (mmGetLocalNodeReference *NodeNetworkMock) GetLocalNodeReference() (h1 reference.Holder) {
	mm_atomic.AddUint64(&mmGetLocalNodeReference.beforeGetLocalNodeReferenceCounter, 1)
	defer mm_atomic.AddUint64(&mmGetLocalNodeReference.afterGetLocalNodeReferenceCounter, 1)

	if mmGetLocalNodeReference.inspectFuncGetLocalNodeReference != nil {
		mmGetLocalNodeReference.inspectFuncGetLocalNodeReference()
	}

	if mmGetLocalNodeReference.GetLocalNodeReferenceMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetLocalNodeReference.GetLocalNodeReferenceMock.defaultExpectation.Counter, 1)

		mm_results := mmGetLocalNodeReference.GetLocalNodeReferenceMock.defaultExpectation.results
		if mm_results == nil {
			mmGetLocalNodeReference.t.Fatal("No results are set for the NodeNetworkMock.GetLocalNodeReference")
		}
		return (*mm_results).h1
	}
	if mmGetLocalNodeReference.funcGetLocalNodeReference != nil {
		return mmGetLocalNodeReference.funcGetLocalNodeReference()
	}
	mmGetLocalNodeReference.t.Fatalf("Unexpected call to NodeNetworkMock.GetLocalNodeReference.")
	return
}

// GetLocalNodeReferenceAfterCounter returns a count of finished NodeNetworkMock.GetLocalNodeReference invocations
func (mmGetLocalNodeReference *NodeNetworkMock) GetLocalNodeReferenceAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetLocalNodeReference.afterGetLocalNodeReferenceCounter)
}

// GetLocalNodeReferenceBeforeCounter returns a count of NodeNetworkMock.GetLocalNodeReference invocations
func (mmGetLocalNodeReference *NodeNetworkMock) GetLocalNodeReferenceBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetLocalNodeReference.beforeGetLocalNodeReferenceCounter)
}

// MinimockGetLocalNodeReferenceDone returns true if the count of the GetLocalNodeReference invocations corresponds
// the number of defined expectations
func (m *NodeNetworkMock) MinimockGetLocalNodeReferenceDone() bool {
	for _, e := range m.GetLocalNodeReferenceMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetLocalNodeReferenceMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetLocalNodeReferenceCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetLocalNodeReference != nil && mm_atomic.LoadUint64(&m.afterGetLocalNodeReferenceCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetLocalNodeReferenceInspect logs each unmet expectation
func (m *NodeNetworkMock) MinimockGetLocalNodeReferenceInspect() {
	for _, e := range m.GetLocalNodeReferenceMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to NodeNetworkMock.GetLocalNodeReference")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetLocalNodeReferenceMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetLocalNodeReferenceCounter) < 1 {
		m.t.Error("Expected call to NodeNetworkMock.GetLocalNodeReference")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetLocalNodeReference != nil && mm_atomic.LoadUint64(&m.afterGetLocalNodeReferenceCounter) < 1 {
		m.t.Error("Expected call to NodeNetworkMock.GetLocalNodeReference")
	}
}

type mNodeNetworkMockGetLocalNodeRole struct {
	mock               *NodeNetworkMock
	defaultExpectation *NodeNetworkMockGetLocalNodeRoleExpectation
	expectations       []*NodeNetworkMockGetLocalNodeRoleExpectation
}

// NodeNetworkMockGetLocalNodeRoleExpectation specifies expectation struct of the NodeNetwork.GetLocalNodeRole
type NodeNetworkMockGetLocalNodeRoleExpectation struct {
	mock *NodeNetworkMock

	results *NodeNetworkMockGetLocalNodeRoleResults
	Counter uint64
}

// NodeNetworkMockGetLocalNodeRoleResults contains results of the NodeNetwork.GetLocalNodeRole
type NodeNetworkMockGetLocalNodeRoleResults struct {
	p1 member.PrimaryRole
}

// Expect sets up expected params for NodeNetwork.GetLocalNodeRole
func (mmGetLocalNodeRole *mNodeNetworkMockGetLocalNodeRole) Expect() *mNodeNetworkMockGetLocalNodeRole {
	if mmGetLocalNodeRole.mock.funcGetLocalNodeRole != nil {
		mmGetLocalNodeRole.mock.t.Fatalf("NodeNetworkMock.GetLocalNodeRole mock is already set by Set")
	}

	if mmGetLocalNodeRole.defaultExpectation == nil {
		mmGetLocalNodeRole.defaultExpectation = &NodeNetworkMockGetLocalNodeRoleExpectation{}
	}

	return mmGetLocalNodeRole
}

// Inspect accepts an inspector function that has same arguments as the NodeNetwork.GetLocalNodeRole
func (mmGetLocalNodeRole *mNodeNetworkMockGetLocalNodeRole) Inspect(f func()) *mNodeNetworkMockGetLocalNodeRole {
	if mmGetLocalNodeRole.mock.inspectFuncGetLocalNodeRole != nil {
		mmGetLocalNodeRole.mock.t.Fatalf("Inspect function is already set for NodeNetworkMock.GetLocalNodeRole")
	}

	mmGetLocalNodeRole.mock.inspectFuncGetLocalNodeRole = f

	return mmGetLocalNodeRole
}

// Return sets up results that will be returned by NodeNetwork.GetLocalNodeRole
func (mmGetLocalNodeRole *mNodeNetworkMockGetLocalNodeRole) Return(p1 member.PrimaryRole) *NodeNetworkMock {
	if mmGetLocalNodeRole.mock.funcGetLocalNodeRole != nil {
		mmGetLocalNodeRole.mock.t.Fatalf("NodeNetworkMock.GetLocalNodeRole mock is already set by Set")
	}

	if mmGetLocalNodeRole.defaultExpectation == nil {
		mmGetLocalNodeRole.defaultExpectation = &NodeNetworkMockGetLocalNodeRoleExpectation{mock: mmGetLocalNodeRole.mock}
	}
	mmGetLocalNodeRole.defaultExpectation.results = &NodeNetworkMockGetLocalNodeRoleResults{p1}
	return mmGetLocalNodeRole.mock
}

//Set uses given function f to mock the NodeNetwork.GetLocalNodeRole method
func (mmGetLocalNodeRole *mNodeNetworkMockGetLocalNodeRole) Set(f func() (p1 member.PrimaryRole)) *NodeNetworkMock {
	if mmGetLocalNodeRole.defaultExpectation != nil {
		mmGetLocalNodeRole.mock.t.Fatalf("Default expectation is already set for the NodeNetwork.GetLocalNodeRole method")
	}

	if len(mmGetLocalNodeRole.expectations) > 0 {
		mmGetLocalNodeRole.mock.t.Fatalf("Some expectations are already set for the NodeNetwork.GetLocalNodeRole method")
	}

	mmGetLocalNodeRole.mock.funcGetLocalNodeRole = f
	return mmGetLocalNodeRole.mock
}

// GetLocalNodeRole implements NodeNetwork
func (mmGetLocalNodeRole *NodeNetworkMock) GetLocalNodeRole() (p1 member.PrimaryRole) {
	mm_atomic.AddUint64(&mmGetLocalNodeRole.beforeGetLocalNodeRoleCounter, 1)
	defer mm_atomic.AddUint64(&mmGetLocalNodeRole.afterGetLocalNodeRoleCounter, 1)

	if mmGetLocalNodeRole.inspectFuncGetLocalNodeRole != nil {
		mmGetLocalNodeRole.inspectFuncGetLocalNodeRole()
	}

	if mmGetLocalNodeRole.GetLocalNodeRoleMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetLocalNodeRole.GetLocalNodeRoleMock.defaultExpectation.Counter, 1)

		mm_results := mmGetLocalNodeRole.GetLocalNodeRoleMock.defaultExpectation.results
		if mm_results == nil {
			mmGetLocalNodeRole.t.Fatal("No results are set for the NodeNetworkMock.GetLocalNodeRole")
		}
		return (*mm_results).p1
	}
	if mmGetLocalNodeRole.funcGetLocalNodeRole != nil {
		return mmGetLocalNodeRole.funcGetLocalNodeRole()
	}
	mmGetLocalNodeRole.t.Fatalf("Unexpected call to NodeNetworkMock.GetLocalNodeRole.")
	return
}

// GetLocalNodeRoleAfterCounter returns a count of finished NodeNetworkMock.GetLocalNodeRole invocations
func (mmGetLocalNodeRole *NodeNetworkMock) GetLocalNodeRoleAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetLocalNodeRole.afterGetLocalNodeRoleCounter)
}

// GetLocalNodeRoleBeforeCounter returns a count of NodeNetworkMock.GetLocalNodeRole invocations
func (mmGetLocalNodeRole *NodeNetworkMock) GetLocalNodeRoleBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetLocalNodeRole.beforeGetLocalNodeRoleCounter)
}

// MinimockGetLocalNodeRoleDone returns true if the count of the GetLocalNodeRole invocations corresponds
// the number of defined expectations
func (m *NodeNetworkMock) MinimockGetLocalNodeRoleDone() bool {
	for _, e := range m.GetLocalNodeRoleMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetLocalNodeRoleMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetLocalNodeRoleCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetLocalNodeRole != nil && mm_atomic.LoadUint64(&m.afterGetLocalNodeRoleCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetLocalNodeRoleInspect logs each unmet expectation
func (m *NodeNetworkMock) MinimockGetLocalNodeRoleInspect() {
	for _, e := range m.GetLocalNodeRoleMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to NodeNetworkMock.GetLocalNodeRole")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetLocalNodeRoleMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetLocalNodeRoleCounter) < 1 {
		m.t.Error("Expected call to NodeNetworkMock.GetLocalNodeRole")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetLocalNodeRole != nil && mm_atomic.LoadUint64(&m.afterGetLocalNodeRoleCounter) < 1 {
		m.t.Error("Expected call to NodeNetworkMock.GetLocalNodeRole")
	}
}

type mNodeNetworkMockGetNodeSnapshot struct {
	mock               *NodeNetworkMock
	defaultExpectation *NodeNetworkMockGetNodeSnapshotExpectation
	expectations       []*NodeNetworkMockGetNodeSnapshotExpectation

	callArgs []*NodeNetworkMockGetNodeSnapshotParams
	mutex    sync.RWMutex
}

// NodeNetworkMockGetNodeSnapshotExpectation specifies expectation struct of the NodeNetwork.GetNodeSnapshot
type NodeNetworkMockGetNodeSnapshotExpectation struct {
	mock    *NodeNetworkMock
	params  *NodeNetworkMockGetNodeSnapshotParams
	results *NodeNetworkMockGetNodeSnapshotResults
	Counter uint64
}

// NodeNetworkMockGetNodeSnapshotParams contains parameters of the NodeNetwork.GetNodeSnapshot
type NodeNetworkMockGetNodeSnapshotParams struct {
	n1 pulse.Number
}

// NodeNetworkMockGetNodeSnapshotResults contains results of the NodeNetwork.GetNodeSnapshot
type NodeNetworkMockGetNodeSnapshotResults struct {
	n2 NodeSnapshot
}

// Expect sets up expected params for NodeNetwork.GetNodeSnapshot
func (mmGetNodeSnapshot *mNodeNetworkMockGetNodeSnapshot) Expect(n1 pulse.Number) *mNodeNetworkMockGetNodeSnapshot {
	if mmGetNodeSnapshot.mock.funcGetNodeSnapshot != nil {
		mmGetNodeSnapshot.mock.t.Fatalf("NodeNetworkMock.GetNodeSnapshot mock is already set by Set")
	}

	if mmGetNodeSnapshot.defaultExpectation == nil {
		mmGetNodeSnapshot.defaultExpectation = &NodeNetworkMockGetNodeSnapshotExpectation{}
	}

	mmGetNodeSnapshot.defaultExpectation.params = &NodeNetworkMockGetNodeSnapshotParams{n1}
	for _, e := range mmGetNodeSnapshot.expectations {
		if minimock.Equal(e.params, mmGetNodeSnapshot.defaultExpectation.params) {
			mmGetNodeSnapshot.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetNodeSnapshot.defaultExpectation.params)
		}
	}

	return mmGetNodeSnapshot
}

// Inspect accepts an inspector function that has same arguments as the NodeNetwork.GetNodeSnapshot
func (mmGetNodeSnapshot *mNodeNetworkMockGetNodeSnapshot) Inspect(f func(n1 pulse.Number)) *mNodeNetworkMockGetNodeSnapshot {
	if mmGetNodeSnapshot.mock.inspectFuncGetNodeSnapshot != nil {
		mmGetNodeSnapshot.mock.t.Fatalf("Inspect function is already set for NodeNetworkMock.GetNodeSnapshot")
	}

	mmGetNodeSnapshot.mock.inspectFuncGetNodeSnapshot = f

	return mmGetNodeSnapshot
}

// Return sets up results that will be returned by NodeNetwork.GetNodeSnapshot
func (mmGetNodeSnapshot *mNodeNetworkMockGetNodeSnapshot) Return(n2 NodeSnapshot) *NodeNetworkMock {
	if mmGetNodeSnapshot.mock.funcGetNodeSnapshot != nil {
		mmGetNodeSnapshot.mock.t.Fatalf("NodeNetworkMock.GetNodeSnapshot mock is already set by Set")
	}

	if mmGetNodeSnapshot.defaultExpectation == nil {
		mmGetNodeSnapshot.defaultExpectation = &NodeNetworkMockGetNodeSnapshotExpectation{mock: mmGetNodeSnapshot.mock}
	}
	mmGetNodeSnapshot.defaultExpectation.results = &NodeNetworkMockGetNodeSnapshotResults{n2}
	return mmGetNodeSnapshot.mock
}

//Set uses given function f to mock the NodeNetwork.GetNodeSnapshot method
func (mmGetNodeSnapshot *mNodeNetworkMockGetNodeSnapshot) Set(f func(n1 pulse.Number) (n2 NodeSnapshot)) *NodeNetworkMock {
	if mmGetNodeSnapshot.defaultExpectation != nil {
		mmGetNodeSnapshot.mock.t.Fatalf("Default expectation is already set for the NodeNetwork.GetNodeSnapshot method")
	}

	if len(mmGetNodeSnapshot.expectations) > 0 {
		mmGetNodeSnapshot.mock.t.Fatalf("Some expectations are already set for the NodeNetwork.GetNodeSnapshot method")
	}

	mmGetNodeSnapshot.mock.funcGetNodeSnapshot = f
	return mmGetNodeSnapshot.mock
}

// When sets expectation for the NodeNetwork.GetNodeSnapshot which will trigger the result defined by the following
// Then helper
func (mmGetNodeSnapshot *mNodeNetworkMockGetNodeSnapshot) When(n1 pulse.Number) *NodeNetworkMockGetNodeSnapshotExpectation {
	if mmGetNodeSnapshot.mock.funcGetNodeSnapshot != nil {
		mmGetNodeSnapshot.mock.t.Fatalf("NodeNetworkMock.GetNodeSnapshot mock is already set by Set")
	}

	expectation := &NodeNetworkMockGetNodeSnapshotExpectation{
		mock:   mmGetNodeSnapshot.mock,
		params: &NodeNetworkMockGetNodeSnapshotParams{n1},
	}
	mmGetNodeSnapshot.expectations = append(mmGetNodeSnapshot.expectations, expectation)
	return expectation
}

// Then sets up NodeNetwork.GetNodeSnapshot return parameters for the expectation previously defined by the When method
func (e *NodeNetworkMockGetNodeSnapshotExpectation) Then(n2 NodeSnapshot) *NodeNetworkMock {
	e.results = &NodeNetworkMockGetNodeSnapshotResults{n2}
	return e.mock
}

// GetNodeSnapshot implements NodeNetwork
func (mmGetNodeSnapshot *NodeNetworkMock) GetNodeSnapshot(n1 pulse.Number) (n2 NodeSnapshot) {
	mm_atomic.AddUint64(&mmGetNodeSnapshot.beforeGetNodeSnapshotCounter, 1)
	defer mm_atomic.AddUint64(&mmGetNodeSnapshot.afterGetNodeSnapshotCounter, 1)

	if mmGetNodeSnapshot.inspectFuncGetNodeSnapshot != nil {
		mmGetNodeSnapshot.inspectFuncGetNodeSnapshot(n1)
	}

	mm_params := &NodeNetworkMockGetNodeSnapshotParams{n1}

	// Record call args
	mmGetNodeSnapshot.GetNodeSnapshotMock.mutex.Lock()
	mmGetNodeSnapshot.GetNodeSnapshotMock.callArgs = append(mmGetNodeSnapshot.GetNodeSnapshotMock.callArgs, mm_params)
	mmGetNodeSnapshot.GetNodeSnapshotMock.mutex.Unlock()

	for _, e := range mmGetNodeSnapshot.GetNodeSnapshotMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.n2
		}
	}

	if mmGetNodeSnapshot.GetNodeSnapshotMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetNodeSnapshot.GetNodeSnapshotMock.defaultExpectation.Counter, 1)
		mm_want := mmGetNodeSnapshot.GetNodeSnapshotMock.defaultExpectation.params
		mm_got := NodeNetworkMockGetNodeSnapshotParams{n1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetNodeSnapshot.t.Errorf("NodeNetworkMock.GetNodeSnapshot got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetNodeSnapshot.GetNodeSnapshotMock.defaultExpectation.results
		if mm_results == nil {
			mmGetNodeSnapshot.t.Fatal("No results are set for the NodeNetworkMock.GetNodeSnapshot")
		}
		return (*mm_results).n2
	}
	if mmGetNodeSnapshot.funcGetNodeSnapshot != nil {
		return mmGetNodeSnapshot.funcGetNodeSnapshot(n1)
	}
	mmGetNodeSnapshot.t.Fatalf("Unexpected call to NodeNetworkMock.GetNodeSnapshot. %v", n1)
	return
}

// GetNodeSnapshotAfterCounter returns a count of finished NodeNetworkMock.GetNodeSnapshot invocations
func (mmGetNodeSnapshot *NodeNetworkMock) GetNodeSnapshotAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetNodeSnapshot.afterGetNodeSnapshotCounter)
}

// GetNodeSnapshotBeforeCounter returns a count of NodeNetworkMock.GetNodeSnapshot invocations
func (mmGetNodeSnapshot *NodeNetworkMock) GetNodeSnapshotBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetNodeSnapshot.beforeGetNodeSnapshotCounter)
}

// Calls returns a list of arguments used in each call to NodeNetworkMock.GetNodeSnapshot.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetNodeSnapshot *mNodeNetworkMockGetNodeSnapshot) Calls() []*NodeNetworkMockGetNodeSnapshotParams {
	mmGetNodeSnapshot.mutex.RLock()

	argCopy := make([]*NodeNetworkMockGetNodeSnapshotParams, len(mmGetNodeSnapshot.callArgs))
	copy(argCopy, mmGetNodeSnapshot.callArgs)

	mmGetNodeSnapshot.mutex.RUnlock()

	return argCopy
}

// MinimockGetNodeSnapshotDone returns true if the count of the GetNodeSnapshot invocations corresponds
// the number of defined expectations
func (m *NodeNetworkMock) MinimockGetNodeSnapshotDone() bool {
	for _, e := range m.GetNodeSnapshotMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetNodeSnapshotMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetNodeSnapshotCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetNodeSnapshot != nil && mm_atomic.LoadUint64(&m.afterGetNodeSnapshotCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetNodeSnapshotInspect logs each unmet expectation
func (m *NodeNetworkMock) MinimockGetNodeSnapshotInspect() {
	for _, e := range m.GetNodeSnapshotMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to NodeNetworkMock.GetNodeSnapshot with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetNodeSnapshotMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetNodeSnapshotCounter) < 1 {
		if m.GetNodeSnapshotMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to NodeNetworkMock.GetNodeSnapshot")
		} else {
			m.t.Errorf("Expected call to NodeNetworkMock.GetNodeSnapshot with params: %#v", *m.GetNodeSnapshotMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetNodeSnapshot != nil && mm_atomic.LoadUint64(&m.afterGetNodeSnapshotCounter) < 1 {
		m.t.Error("Expected call to NodeNetworkMock.GetNodeSnapshot")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *NodeNetworkMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetAnyLatestNodeSnapshotInspect()

		m.MinimockGetLocalNodeReferenceInspect()

		m.MinimockGetLocalNodeRoleInspect()

		m.MinimockGetNodeSnapshotInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *NodeNetworkMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *NodeNetworkMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetAnyLatestNodeSnapshotDone() &&
		m.MinimockGetLocalNodeReferenceDone() &&
		m.MinimockGetLocalNodeRoleDone() &&
		m.MinimockGetNodeSnapshotDone()
}
