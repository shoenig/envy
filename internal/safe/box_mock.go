package safe

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/hashicorp/go-set/v2"
)

// BoxMock implements Box
type BoxMock struct {
	t minimock.Tester

	funcDelete          func(s1 string, pp1 *set.Set[string]) (err error)
	inspectFuncDelete   func(s1 string, pp1 *set.Set[string])
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mBoxMockDelete

	funcGet          func(s1 string) (np1 *Namespace, err error)
	inspectFuncGet   func(s1 string)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mBoxMockGet

	funcList          func() (sa1 []string, err error)
	inspectFuncList   func()
	afterListCounter  uint64
	beforeListCounter uint64
	ListMock          mBoxMockList

	funcPurge          func(s1 string) (err error)
	inspectFuncPurge   func(s1 string)
	afterPurgeCounter  uint64
	beforePurgeCounter uint64
	PurgeMock          mBoxMockPurge

	funcSet          func(np1 *Namespace) (err error)
	inspectFuncSet   func(np1 *Namespace)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mBoxMockSet
}

// NewBoxMock returns a mock for Box
func NewBoxMock(t minimock.Tester) *BoxMock {
	m := &BoxMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeleteMock = mBoxMockDelete{mock: m}
	m.DeleteMock.callArgs = []*BoxMockDeleteParams{}

	m.GetMock = mBoxMockGet{mock: m}
	m.GetMock.callArgs = []*BoxMockGetParams{}

	m.ListMock = mBoxMockList{mock: m}

	m.PurgeMock = mBoxMockPurge{mock: m}
	m.PurgeMock.callArgs = []*BoxMockPurgeParams{}

	m.SetMock = mBoxMockSet{mock: m}
	m.SetMock.callArgs = []*BoxMockSetParams{}

	return m
}

type mBoxMockDelete struct {
	mock               *BoxMock
	defaultExpectation *BoxMockDeleteExpectation
	expectations       []*BoxMockDeleteExpectation

	callArgs []*BoxMockDeleteParams
	mutex    sync.RWMutex
}

// BoxMockDeleteExpectation specifies expectation struct of the Box.Delete
type BoxMockDeleteExpectation struct {
	mock    *BoxMock
	params  *BoxMockDeleteParams
	results *BoxMockDeleteResults
	Counter uint64
}

// BoxMockDeleteParams contains parameters of the Box.Delete
type BoxMockDeleteParams struct {
	s1  string
	pp1 *set.Set[string]
}

// BoxMockDeleteResults contains results of the Box.Delete
type BoxMockDeleteResults struct {
	err error
}

// Expect sets up expected params for Box.Delete
func (mmDelete *mBoxMockDelete) Expect(s1 string, pp1 *set.Set[string]) *mBoxMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("BoxMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &BoxMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &BoxMockDeleteParams{s1, pp1}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the Box.Delete
func (mmDelete *mBoxMockDelete) Inspect(f func(s1 string, pp1 *set.Set[string])) *mBoxMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for BoxMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by Box.Delete
func (mmDelete *mBoxMockDelete) Return(err error) *BoxMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("BoxMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &BoxMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &BoxMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the Box.Delete method
func (mmDelete *mBoxMockDelete) Set(f func(s1 string, pp1 *set.Set[string]) (err error)) *BoxMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the Box.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the Box.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the Box.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mBoxMockDelete) When(s1 string, pp1 *set.Set[string]) *BoxMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("BoxMock.Delete mock is already set by Set")
	}

	expectation := &BoxMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &BoxMockDeleteParams{s1, pp1},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up Box.Delete return parameters for the expectation previously defined by the When method
func (e *BoxMockDeleteExpectation) Then(err error) *BoxMock {
	e.results = &BoxMockDeleteResults{err}
	return e.mock
}

// Delete implements Box
func (mmDelete *BoxMock) Delete(s1 string, pp1 *set.Set[string]) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(s1, pp1)
	}

	mm_params := &BoxMockDeleteParams{s1, pp1}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := BoxMockDeleteParams{s1, pp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("BoxMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the BoxMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(s1, pp1)
	}
	mmDelete.t.Fatalf("Unexpected call to BoxMock.Delete. %v %v", s1, pp1)
	return
}

// DeleteAfterCounter returns a count of finished BoxMock.Delete invocations
func (mmDelete *BoxMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of BoxMock.Delete invocations
func (mmDelete *BoxMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to BoxMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mBoxMockDelete) Calls() []*BoxMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*BoxMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *BoxMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *BoxMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BoxMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BoxMock.Delete")
		} else {
			m.t.Errorf("Expected call to BoxMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to BoxMock.Delete")
	}
}

type mBoxMockGet struct {
	mock               *BoxMock
	defaultExpectation *BoxMockGetExpectation
	expectations       []*BoxMockGetExpectation

	callArgs []*BoxMockGetParams
	mutex    sync.RWMutex
}

// BoxMockGetExpectation specifies expectation struct of the Box.Get
type BoxMockGetExpectation struct {
	mock    *BoxMock
	params  *BoxMockGetParams
	results *BoxMockGetResults
	Counter uint64
}

// BoxMockGetParams contains parameters of the Box.Get
type BoxMockGetParams struct {
	s1 string
}

// BoxMockGetResults contains results of the Box.Get
type BoxMockGetResults struct {
	np1 *Namespace
	err error
}

// Expect sets up expected params for Box.Get
func (mmGet *mBoxMockGet) Expect(s1 string) *mBoxMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("BoxMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &BoxMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &BoxMockGetParams{s1}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the Box.Get
func (mmGet *mBoxMockGet) Inspect(f func(s1 string)) *mBoxMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for BoxMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by Box.Get
func (mmGet *mBoxMockGet) Return(np1 *Namespace, err error) *BoxMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("BoxMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &BoxMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &BoxMockGetResults{np1, err}
	return mmGet.mock
}

// Set uses given function f to mock the Box.Get method
func (mmGet *mBoxMockGet) Set(f func(s1 string) (np1 *Namespace, err error)) *BoxMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the Box.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the Box.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the Box.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mBoxMockGet) When(s1 string) *BoxMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("BoxMock.Get mock is already set by Set")
	}

	expectation := &BoxMockGetExpectation{
		mock:   mmGet.mock,
		params: &BoxMockGetParams{s1},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up Box.Get return parameters for the expectation previously defined by the When method
func (e *BoxMockGetExpectation) Then(np1 *Namespace, err error) *BoxMock {
	e.results = &BoxMockGetResults{np1, err}
	return e.mock
}

// Get implements Box
func (mmGet *BoxMock) Get(s1 string) (np1 *Namespace, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(s1)
	}

	mm_params := &BoxMockGetParams{s1}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.np1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := BoxMockGetParams{s1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("BoxMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the BoxMock.Get")
		}
		return (*mm_results).np1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(s1)
	}
	mmGet.t.Fatalf("Unexpected call to BoxMock.Get. %v", s1)
	return
}

// GetAfterCounter returns a count of finished BoxMock.Get invocations
func (mmGet *BoxMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of BoxMock.Get invocations
func (mmGet *BoxMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to BoxMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mBoxMockGet) Calls() []*BoxMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*BoxMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *BoxMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *BoxMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BoxMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BoxMock.Get")
		} else {
			m.t.Errorf("Expected call to BoxMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to BoxMock.Get")
	}
}

type mBoxMockList struct {
	mock               *BoxMock
	defaultExpectation *BoxMockListExpectation
	expectations       []*BoxMockListExpectation
}

// BoxMockListExpectation specifies expectation struct of the Box.List
type BoxMockListExpectation struct {
	mock *BoxMock

	results *BoxMockListResults
	Counter uint64
}

// BoxMockListResults contains results of the Box.List
type BoxMockListResults struct {
	sa1 []string
	err error
}

// Expect sets up expected params for Box.List
func (mmList *mBoxMockList) Expect() *mBoxMockList {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("BoxMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &BoxMockListExpectation{}
	}

	return mmList
}

// Inspect accepts an inspector function that has same arguments as the Box.List
func (mmList *mBoxMockList) Inspect(f func()) *mBoxMockList {
	if mmList.mock.inspectFuncList != nil {
		mmList.mock.t.Fatalf("Inspect function is already set for BoxMock.List")
	}

	mmList.mock.inspectFuncList = f

	return mmList
}

// Return sets up results that will be returned by Box.List
func (mmList *mBoxMockList) Return(sa1 []string, err error) *BoxMock {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("BoxMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &BoxMockListExpectation{mock: mmList.mock}
	}
	mmList.defaultExpectation.results = &BoxMockListResults{sa1, err}
	return mmList.mock
}

// Set uses given function f to mock the Box.List method
func (mmList *mBoxMockList) Set(f func() (sa1 []string, err error)) *BoxMock {
	if mmList.defaultExpectation != nil {
		mmList.mock.t.Fatalf("Default expectation is already set for the Box.List method")
	}

	if len(mmList.expectations) > 0 {
		mmList.mock.t.Fatalf("Some expectations are already set for the Box.List method")
	}

	mmList.mock.funcList = f
	return mmList.mock
}

// List implements Box
func (mmList *BoxMock) List() (sa1 []string, err error) {
	mm_atomic.AddUint64(&mmList.beforeListCounter, 1)
	defer mm_atomic.AddUint64(&mmList.afterListCounter, 1)

	if mmList.inspectFuncList != nil {
		mmList.inspectFuncList()
	}

	if mmList.ListMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmList.ListMock.defaultExpectation.Counter, 1)

		mm_results := mmList.ListMock.defaultExpectation.results
		if mm_results == nil {
			mmList.t.Fatal("No results are set for the BoxMock.List")
		}
		return (*mm_results).sa1, (*mm_results).err
	}
	if mmList.funcList != nil {
		return mmList.funcList()
	}
	mmList.t.Fatalf("Unexpected call to BoxMock.List.")
	return
}

// ListAfterCounter returns a count of finished BoxMock.List invocations
func (mmList *BoxMock) ListAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmList.afterListCounter)
}

// ListBeforeCounter returns a count of BoxMock.List invocations
func (mmList *BoxMock) ListBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmList.beforeListCounter)
}

// MinimockListDone returns true if the count of the List invocations corresponds
// the number of defined expectations
func (m *BoxMock) MinimockListDone() bool {
	for _, e := range m.ListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcList != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		return false
	}
	return true
}

// MinimockListInspect logs each unmet expectation
func (m *BoxMock) MinimockListInspect() {
	for _, e := range m.ListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to BoxMock.List")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		m.t.Error("Expected call to BoxMock.List")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcList != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		m.t.Error("Expected call to BoxMock.List")
	}
}

type mBoxMockPurge struct {
	mock               *BoxMock
	defaultExpectation *BoxMockPurgeExpectation
	expectations       []*BoxMockPurgeExpectation

	callArgs []*BoxMockPurgeParams
	mutex    sync.RWMutex
}

// BoxMockPurgeExpectation specifies expectation struct of the Box.Purge
type BoxMockPurgeExpectation struct {
	mock    *BoxMock
	params  *BoxMockPurgeParams
	results *BoxMockPurgeResults
	Counter uint64
}

// BoxMockPurgeParams contains parameters of the Box.Purge
type BoxMockPurgeParams struct {
	s1 string
}

// BoxMockPurgeResults contains results of the Box.Purge
type BoxMockPurgeResults struct {
	err error
}

// Expect sets up expected params for Box.Purge
func (mmPurge *mBoxMockPurge) Expect(s1 string) *mBoxMockPurge {
	if mmPurge.mock.funcPurge != nil {
		mmPurge.mock.t.Fatalf("BoxMock.Purge mock is already set by Set")
	}

	if mmPurge.defaultExpectation == nil {
		mmPurge.defaultExpectation = &BoxMockPurgeExpectation{}
	}

	mmPurge.defaultExpectation.params = &BoxMockPurgeParams{s1}
	for _, e := range mmPurge.expectations {
		if minimock.Equal(e.params, mmPurge.defaultExpectation.params) {
			mmPurge.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmPurge.defaultExpectation.params)
		}
	}

	return mmPurge
}

// Inspect accepts an inspector function that has same arguments as the Box.Purge
func (mmPurge *mBoxMockPurge) Inspect(f func(s1 string)) *mBoxMockPurge {
	if mmPurge.mock.inspectFuncPurge != nil {
		mmPurge.mock.t.Fatalf("Inspect function is already set for BoxMock.Purge")
	}

	mmPurge.mock.inspectFuncPurge = f

	return mmPurge
}

// Return sets up results that will be returned by Box.Purge
func (mmPurge *mBoxMockPurge) Return(err error) *BoxMock {
	if mmPurge.mock.funcPurge != nil {
		mmPurge.mock.t.Fatalf("BoxMock.Purge mock is already set by Set")
	}

	if mmPurge.defaultExpectation == nil {
		mmPurge.defaultExpectation = &BoxMockPurgeExpectation{mock: mmPurge.mock}
	}
	mmPurge.defaultExpectation.results = &BoxMockPurgeResults{err}
	return mmPurge.mock
}

// Set uses given function f to mock the Box.Purge method
func (mmPurge *mBoxMockPurge) Set(f func(s1 string) (err error)) *BoxMock {
	if mmPurge.defaultExpectation != nil {
		mmPurge.mock.t.Fatalf("Default expectation is already set for the Box.Purge method")
	}

	if len(mmPurge.expectations) > 0 {
		mmPurge.mock.t.Fatalf("Some expectations are already set for the Box.Purge method")
	}

	mmPurge.mock.funcPurge = f
	return mmPurge.mock
}

// When sets expectation for the Box.Purge which will trigger the result defined by the following
// Then helper
func (mmPurge *mBoxMockPurge) When(s1 string) *BoxMockPurgeExpectation {
	if mmPurge.mock.funcPurge != nil {
		mmPurge.mock.t.Fatalf("BoxMock.Purge mock is already set by Set")
	}

	expectation := &BoxMockPurgeExpectation{
		mock:   mmPurge.mock,
		params: &BoxMockPurgeParams{s1},
	}
	mmPurge.expectations = append(mmPurge.expectations, expectation)
	return expectation
}

// Then sets up Box.Purge return parameters for the expectation previously defined by the When method
func (e *BoxMockPurgeExpectation) Then(err error) *BoxMock {
	e.results = &BoxMockPurgeResults{err}
	return e.mock
}

// Purge implements Box
func (mmPurge *BoxMock) Purge(s1 string) (err error) {
	mm_atomic.AddUint64(&mmPurge.beforePurgeCounter, 1)
	defer mm_atomic.AddUint64(&mmPurge.afterPurgeCounter, 1)

	if mmPurge.inspectFuncPurge != nil {
		mmPurge.inspectFuncPurge(s1)
	}

	mm_params := &BoxMockPurgeParams{s1}

	// Record call args
	mmPurge.PurgeMock.mutex.Lock()
	mmPurge.PurgeMock.callArgs = append(mmPurge.PurgeMock.callArgs, mm_params)
	mmPurge.PurgeMock.mutex.Unlock()

	for _, e := range mmPurge.PurgeMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmPurge.PurgeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPurge.PurgeMock.defaultExpectation.Counter, 1)
		mm_want := mmPurge.PurgeMock.defaultExpectation.params
		mm_got := BoxMockPurgeParams{s1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmPurge.t.Errorf("BoxMock.Purge got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmPurge.PurgeMock.defaultExpectation.results
		if mm_results == nil {
			mmPurge.t.Fatal("No results are set for the BoxMock.Purge")
		}
		return (*mm_results).err
	}
	if mmPurge.funcPurge != nil {
		return mmPurge.funcPurge(s1)
	}
	mmPurge.t.Fatalf("Unexpected call to BoxMock.Purge. %v", s1)
	return
}

// PurgeAfterCounter returns a count of finished BoxMock.Purge invocations
func (mmPurge *BoxMock) PurgeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPurge.afterPurgeCounter)
}

// PurgeBeforeCounter returns a count of BoxMock.Purge invocations
func (mmPurge *BoxMock) PurgeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPurge.beforePurgeCounter)
}

// Calls returns a list of arguments used in each call to BoxMock.Purge.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmPurge *mBoxMockPurge) Calls() []*BoxMockPurgeParams {
	mmPurge.mutex.RLock()

	argCopy := make([]*BoxMockPurgeParams, len(mmPurge.callArgs))
	copy(argCopy, mmPurge.callArgs)

	mmPurge.mutex.RUnlock()

	return argCopy
}

// MinimockPurgeDone returns true if the count of the Purge invocations corresponds
// the number of defined expectations
func (m *BoxMock) MinimockPurgeDone() bool {
	for _, e := range m.PurgeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PurgeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPurgeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPurge != nil && mm_atomic.LoadUint64(&m.afterPurgeCounter) < 1 {
		return false
	}
	return true
}

// MinimockPurgeInspect logs each unmet expectation
func (m *BoxMock) MinimockPurgeInspect() {
	for _, e := range m.PurgeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BoxMock.Purge with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PurgeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPurgeCounter) < 1 {
		if m.PurgeMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BoxMock.Purge")
		} else {
			m.t.Errorf("Expected call to BoxMock.Purge with params: %#v", *m.PurgeMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPurge != nil && mm_atomic.LoadUint64(&m.afterPurgeCounter) < 1 {
		m.t.Error("Expected call to BoxMock.Purge")
	}
}

type mBoxMockSet struct {
	mock               *BoxMock
	defaultExpectation *BoxMockSetExpectation
	expectations       []*BoxMockSetExpectation

	callArgs []*BoxMockSetParams
	mutex    sync.RWMutex
}

// BoxMockSetExpectation specifies expectation struct of the Box.Set
type BoxMockSetExpectation struct {
	mock    *BoxMock
	params  *BoxMockSetParams
	results *BoxMockSetResults
	Counter uint64
}

// BoxMockSetParams contains parameters of the Box.Set
type BoxMockSetParams struct {
	np1 *Namespace
}

// BoxMockSetResults contains results of the Box.Set
type BoxMockSetResults struct {
	err error
}

// Expect sets up expected params for Box.Set
func (mmSet *mBoxMockSet) Expect(np1 *Namespace) *mBoxMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("BoxMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &BoxMockSetExpectation{}
	}

	mmSet.defaultExpectation.params = &BoxMockSetParams{np1}
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// Inspect accepts an inspector function that has same arguments as the Box.Set
func (mmSet *mBoxMockSet) Inspect(f func(np1 *Namespace)) *mBoxMockSet {
	if mmSet.mock.inspectFuncSet != nil {
		mmSet.mock.t.Fatalf("Inspect function is already set for BoxMock.Set")
	}

	mmSet.mock.inspectFuncSet = f

	return mmSet
}

// Return sets up results that will be returned by Box.Set
func (mmSet *mBoxMockSet) Return(err error) *BoxMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("BoxMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &BoxMockSetExpectation{mock: mmSet.mock}
	}
	mmSet.defaultExpectation.results = &BoxMockSetResults{err}
	return mmSet.mock
}

// Set uses given function f to mock the Box.Set method
func (mmSet *mBoxMockSet) Set(f func(np1 *Namespace) (err error)) *BoxMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the Box.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the Box.Set method")
	}

	mmSet.mock.funcSet = f
	return mmSet.mock
}

// When sets expectation for the Box.Set which will trigger the result defined by the following
// Then helper
func (mmSet *mBoxMockSet) When(np1 *Namespace) *BoxMockSetExpectation {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("BoxMock.Set mock is already set by Set")
	}

	expectation := &BoxMockSetExpectation{
		mock:   mmSet.mock,
		params: &BoxMockSetParams{np1},
	}
	mmSet.expectations = append(mmSet.expectations, expectation)
	return expectation
}

// Then sets up Box.Set return parameters for the expectation previously defined by the When method
func (e *BoxMockSetExpectation) Then(err error) *BoxMock {
	e.results = &BoxMockSetResults{err}
	return e.mock
}

// Set implements Box
func (mmSet *BoxMock) Set(np1 *Namespace) (err error) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	if mmSet.inspectFuncSet != nil {
		mmSet.inspectFuncSet(np1)
	}

	mm_params := &BoxMockSetParams{np1}

	// Record call args
	mmSet.SetMock.mutex.Lock()
	mmSet.SetMock.callArgs = append(mmSet.SetMock.callArgs, mm_params)
	mmSet.SetMock.mutex.Unlock()

	for _, e := range mmSet.SetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSet.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSet.SetMock.defaultExpectation.Counter, 1)
		mm_want := mmSet.SetMock.defaultExpectation.params
		mm_got := BoxMockSetParams{np1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSet.t.Errorf("BoxMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSet.SetMock.defaultExpectation.results
		if mm_results == nil {
			mmSet.t.Fatal("No results are set for the BoxMock.Set")
		}
		return (*mm_results).err
	}
	if mmSet.funcSet != nil {
		return mmSet.funcSet(np1)
	}
	mmSet.t.Fatalf("Unexpected call to BoxMock.Set. %v", np1)
	return
}

// SetAfterCounter returns a count of finished BoxMock.Set invocations
func (mmSet *BoxMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of BoxMock.Set invocations
func (mmSet *BoxMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to BoxMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mBoxMockSet) Calls() []*BoxMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*BoxMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *BoxMock) MinimockSetDone() bool {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetInspect logs each unmet expectation
func (m *BoxMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BoxMock.Set with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BoxMock.Set")
		} else {
			m.t.Errorf("Expected call to BoxMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		m.t.Error("Expected call to BoxMock.Set")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *BoxMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDeleteInspect()

		m.MinimockGetInspect()

		m.MinimockListInspect()

		m.MinimockPurgeInspect()

		m.MinimockSetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *BoxMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *BoxMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDeleteDone() &&
		m.MinimockGetDone() &&
		m.MinimockListDone() &&
		m.MinimockPurgeDone() &&
		m.MinimockSetDone()
}
