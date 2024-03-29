package keyring

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/go-conceal"
)

// RingMock implements Ring
type RingMock struct {
	t minimock.Tester

	funcDecrypt          func(e1 safe.Encrypted) (tp1 *conceal.Text)
	inspectFuncDecrypt   func(e1 safe.Encrypted)
	afterDecryptCounter  uint64
	beforeDecryptCounter uint64
	DecryptMock          mRingMockDecrypt

	funcEncrypt          func(tp1 *conceal.Text) (e1 safe.Encrypted)
	inspectFuncEncrypt   func(tp1 *conceal.Text)
	afterEncryptCounter  uint64
	beforeEncryptCounter uint64
	EncryptMock          mRingMockEncrypt
}

// NewRingMock returns a mock for Ring
func NewRingMock(t minimock.Tester) *RingMock {
	m := &RingMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DecryptMock = mRingMockDecrypt{mock: m}
	m.DecryptMock.callArgs = []*RingMockDecryptParams{}

	m.EncryptMock = mRingMockEncrypt{mock: m}
	m.EncryptMock.callArgs = []*RingMockEncryptParams{}

	return m
}

type mRingMockDecrypt struct {
	mock               *RingMock
	defaultExpectation *RingMockDecryptExpectation
	expectations       []*RingMockDecryptExpectation

	callArgs []*RingMockDecryptParams
	mutex    sync.RWMutex
}

// RingMockDecryptExpectation specifies expectation struct of the Ring.Decrypt
type RingMockDecryptExpectation struct {
	mock    *RingMock
	params  *RingMockDecryptParams
	results *RingMockDecryptResults
	Counter uint64
}

// RingMockDecryptParams contains parameters of the Ring.Decrypt
type RingMockDecryptParams struct {
	e1 safe.Encrypted
}

// RingMockDecryptResults contains results of the Ring.Decrypt
type RingMockDecryptResults struct {
	tp1 *conceal.Text
}

// Expect sets up expected params for Ring.Decrypt
func (mmDecrypt *mRingMockDecrypt) Expect(e1 safe.Encrypted) *mRingMockDecrypt {
	if mmDecrypt.mock.funcDecrypt != nil {
		mmDecrypt.mock.t.Fatalf("RingMock.Decrypt mock is already set by Set")
	}

	if mmDecrypt.defaultExpectation == nil {
		mmDecrypt.defaultExpectation = &RingMockDecryptExpectation{}
	}

	mmDecrypt.defaultExpectation.params = &RingMockDecryptParams{e1}
	for _, e := range mmDecrypt.expectations {
		if minimock.Equal(e.params, mmDecrypt.defaultExpectation.params) {
			mmDecrypt.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDecrypt.defaultExpectation.params)
		}
	}

	return mmDecrypt
}

// Inspect accepts an inspector function that has same arguments as the Ring.Decrypt
func (mmDecrypt *mRingMockDecrypt) Inspect(f func(e1 safe.Encrypted)) *mRingMockDecrypt {
	if mmDecrypt.mock.inspectFuncDecrypt != nil {
		mmDecrypt.mock.t.Fatalf("Inspect function is already set for RingMock.Decrypt")
	}

	mmDecrypt.mock.inspectFuncDecrypt = f

	return mmDecrypt
}

// Return sets up results that will be returned by Ring.Decrypt
func (mmDecrypt *mRingMockDecrypt) Return(tp1 *conceal.Text) *RingMock {
	if mmDecrypt.mock.funcDecrypt != nil {
		mmDecrypt.mock.t.Fatalf("RingMock.Decrypt mock is already set by Set")
	}

	if mmDecrypt.defaultExpectation == nil {
		mmDecrypt.defaultExpectation = &RingMockDecryptExpectation{mock: mmDecrypt.mock}
	}
	mmDecrypt.defaultExpectation.results = &RingMockDecryptResults{tp1}
	return mmDecrypt.mock
}

// Set uses given function f to mock the Ring.Decrypt method
func (mmDecrypt *mRingMockDecrypt) Set(f func(e1 safe.Encrypted) (tp1 *conceal.Text)) *RingMock {
	if mmDecrypt.defaultExpectation != nil {
		mmDecrypt.mock.t.Fatalf("Default expectation is already set for the Ring.Decrypt method")
	}

	if len(mmDecrypt.expectations) > 0 {
		mmDecrypt.mock.t.Fatalf("Some expectations are already set for the Ring.Decrypt method")
	}

	mmDecrypt.mock.funcDecrypt = f
	return mmDecrypt.mock
}

// When sets expectation for the Ring.Decrypt which will trigger the result defined by the following
// Then helper
func (mmDecrypt *mRingMockDecrypt) When(e1 safe.Encrypted) *RingMockDecryptExpectation {
	if mmDecrypt.mock.funcDecrypt != nil {
		mmDecrypt.mock.t.Fatalf("RingMock.Decrypt mock is already set by Set")
	}

	expectation := &RingMockDecryptExpectation{
		mock:   mmDecrypt.mock,
		params: &RingMockDecryptParams{e1},
	}
	mmDecrypt.expectations = append(mmDecrypt.expectations, expectation)
	return expectation
}

// Then sets up Ring.Decrypt return parameters for the expectation previously defined by the When method
func (e *RingMockDecryptExpectation) Then(tp1 *conceal.Text) *RingMock {
	e.results = &RingMockDecryptResults{tp1}
	return e.mock
}

// Decrypt implements Ring
func (mmDecrypt *RingMock) Decrypt(e1 safe.Encrypted) (tp1 *conceal.Text) {
	mm_atomic.AddUint64(&mmDecrypt.beforeDecryptCounter, 1)
	defer mm_atomic.AddUint64(&mmDecrypt.afterDecryptCounter, 1)

	if mmDecrypt.inspectFuncDecrypt != nil {
		mmDecrypt.inspectFuncDecrypt(e1)
	}

	mm_params := &RingMockDecryptParams{e1}

	// Record call args
	mmDecrypt.DecryptMock.mutex.Lock()
	mmDecrypt.DecryptMock.callArgs = append(mmDecrypt.DecryptMock.callArgs, mm_params)
	mmDecrypt.DecryptMock.mutex.Unlock()

	for _, e := range mmDecrypt.DecryptMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.tp1
		}
	}

	if mmDecrypt.DecryptMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDecrypt.DecryptMock.defaultExpectation.Counter, 1)
		mm_want := mmDecrypt.DecryptMock.defaultExpectation.params
		mm_got := RingMockDecryptParams{e1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDecrypt.t.Errorf("RingMock.Decrypt got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDecrypt.DecryptMock.defaultExpectation.results
		if mm_results == nil {
			mmDecrypt.t.Fatal("No results are set for the RingMock.Decrypt")
		}
		return (*mm_results).tp1
	}
	if mmDecrypt.funcDecrypt != nil {
		return mmDecrypt.funcDecrypt(e1)
	}
	mmDecrypt.t.Fatalf("Unexpected call to RingMock.Decrypt. %v", e1)
	return
}

// DecryptAfterCounter returns a count of finished RingMock.Decrypt invocations
func (mmDecrypt *RingMock) DecryptAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDecrypt.afterDecryptCounter)
}

// DecryptBeforeCounter returns a count of RingMock.Decrypt invocations
func (mmDecrypt *RingMock) DecryptBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDecrypt.beforeDecryptCounter)
}

// Calls returns a list of arguments used in each call to RingMock.Decrypt.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDecrypt *mRingMockDecrypt) Calls() []*RingMockDecryptParams {
	mmDecrypt.mutex.RLock()

	argCopy := make([]*RingMockDecryptParams, len(mmDecrypt.callArgs))
	copy(argCopy, mmDecrypt.callArgs)

	mmDecrypt.mutex.RUnlock()

	return argCopy
}

// MinimockDecryptDone returns true if the count of the Decrypt invocations corresponds
// the number of defined expectations
func (m *RingMock) MinimockDecryptDone() bool {
	for _, e := range m.DecryptMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DecryptMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDecryptCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDecrypt != nil && mm_atomic.LoadUint64(&m.afterDecryptCounter) < 1 {
		return false
	}
	return true
}

// MinimockDecryptInspect logs each unmet expectation
func (m *RingMock) MinimockDecryptInspect() {
	for _, e := range m.DecryptMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RingMock.Decrypt with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DecryptMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDecryptCounter) < 1 {
		if m.DecryptMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RingMock.Decrypt")
		} else {
			m.t.Errorf("Expected call to RingMock.Decrypt with params: %#v", *m.DecryptMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDecrypt != nil && mm_atomic.LoadUint64(&m.afterDecryptCounter) < 1 {
		m.t.Error("Expected call to RingMock.Decrypt")
	}
}

type mRingMockEncrypt struct {
	mock               *RingMock
	defaultExpectation *RingMockEncryptExpectation
	expectations       []*RingMockEncryptExpectation

	callArgs []*RingMockEncryptParams
	mutex    sync.RWMutex
}

// RingMockEncryptExpectation specifies expectation struct of the Ring.Encrypt
type RingMockEncryptExpectation struct {
	mock    *RingMock
	params  *RingMockEncryptParams
	results *RingMockEncryptResults
	Counter uint64
}

// RingMockEncryptParams contains parameters of the Ring.Encrypt
type RingMockEncryptParams struct {
	tp1 *conceal.Text
}

// RingMockEncryptResults contains results of the Ring.Encrypt
type RingMockEncryptResults struct {
	e1 safe.Encrypted
}

// Expect sets up expected params for Ring.Encrypt
func (mmEncrypt *mRingMockEncrypt) Expect(tp1 *conceal.Text) *mRingMockEncrypt {
	if mmEncrypt.mock.funcEncrypt != nil {
		mmEncrypt.mock.t.Fatalf("RingMock.Encrypt mock is already set by Set")
	}

	if mmEncrypt.defaultExpectation == nil {
		mmEncrypt.defaultExpectation = &RingMockEncryptExpectation{}
	}

	mmEncrypt.defaultExpectation.params = &RingMockEncryptParams{tp1}
	for _, e := range mmEncrypt.expectations {
		if minimock.Equal(e.params, mmEncrypt.defaultExpectation.params) {
			mmEncrypt.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmEncrypt.defaultExpectation.params)
		}
	}

	return mmEncrypt
}

// Inspect accepts an inspector function that has same arguments as the Ring.Encrypt
func (mmEncrypt *mRingMockEncrypt) Inspect(f func(tp1 *conceal.Text)) *mRingMockEncrypt {
	if mmEncrypt.mock.inspectFuncEncrypt != nil {
		mmEncrypt.mock.t.Fatalf("Inspect function is already set for RingMock.Encrypt")
	}

	mmEncrypt.mock.inspectFuncEncrypt = f

	return mmEncrypt
}

// Return sets up results that will be returned by Ring.Encrypt
func (mmEncrypt *mRingMockEncrypt) Return(e1 safe.Encrypted) *RingMock {
	if mmEncrypt.mock.funcEncrypt != nil {
		mmEncrypt.mock.t.Fatalf("RingMock.Encrypt mock is already set by Set")
	}

	if mmEncrypt.defaultExpectation == nil {
		mmEncrypt.defaultExpectation = &RingMockEncryptExpectation{mock: mmEncrypt.mock}
	}
	mmEncrypt.defaultExpectation.results = &RingMockEncryptResults{e1}
	return mmEncrypt.mock
}

// Set uses given function f to mock the Ring.Encrypt method
func (mmEncrypt *mRingMockEncrypt) Set(f func(tp1 *conceal.Text) (e1 safe.Encrypted)) *RingMock {
	if mmEncrypt.defaultExpectation != nil {
		mmEncrypt.mock.t.Fatalf("Default expectation is already set for the Ring.Encrypt method")
	}

	if len(mmEncrypt.expectations) > 0 {
		mmEncrypt.mock.t.Fatalf("Some expectations are already set for the Ring.Encrypt method")
	}

	mmEncrypt.mock.funcEncrypt = f
	return mmEncrypt.mock
}

// When sets expectation for the Ring.Encrypt which will trigger the result defined by the following
// Then helper
func (mmEncrypt *mRingMockEncrypt) When(tp1 *conceal.Text) *RingMockEncryptExpectation {
	if mmEncrypt.mock.funcEncrypt != nil {
		mmEncrypt.mock.t.Fatalf("RingMock.Encrypt mock is already set by Set")
	}

	expectation := &RingMockEncryptExpectation{
		mock:   mmEncrypt.mock,
		params: &RingMockEncryptParams{tp1},
	}
	mmEncrypt.expectations = append(mmEncrypt.expectations, expectation)
	return expectation
}

// Then sets up Ring.Encrypt return parameters for the expectation previously defined by the When method
func (e *RingMockEncryptExpectation) Then(e1 safe.Encrypted) *RingMock {
	e.results = &RingMockEncryptResults{e1}
	return e.mock
}

// Encrypt implements Ring
func (mmEncrypt *RingMock) Encrypt(tp1 *conceal.Text) (e1 safe.Encrypted) {
	mm_atomic.AddUint64(&mmEncrypt.beforeEncryptCounter, 1)
	defer mm_atomic.AddUint64(&mmEncrypt.afterEncryptCounter, 1)

	if mmEncrypt.inspectFuncEncrypt != nil {
		mmEncrypt.inspectFuncEncrypt(tp1)
	}

	mm_params := &RingMockEncryptParams{tp1}

	// Record call args
	mmEncrypt.EncryptMock.mutex.Lock()
	mmEncrypt.EncryptMock.callArgs = append(mmEncrypt.EncryptMock.callArgs, mm_params)
	mmEncrypt.EncryptMock.mutex.Unlock()

	for _, e := range mmEncrypt.EncryptMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.e1
		}
	}

	if mmEncrypt.EncryptMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmEncrypt.EncryptMock.defaultExpectation.Counter, 1)
		mm_want := mmEncrypt.EncryptMock.defaultExpectation.params
		mm_got := RingMockEncryptParams{tp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmEncrypt.t.Errorf("RingMock.Encrypt got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmEncrypt.EncryptMock.defaultExpectation.results
		if mm_results == nil {
			mmEncrypt.t.Fatal("No results are set for the RingMock.Encrypt")
		}
		return (*mm_results).e1
	}
	if mmEncrypt.funcEncrypt != nil {
		return mmEncrypt.funcEncrypt(tp1)
	}
	mmEncrypt.t.Fatalf("Unexpected call to RingMock.Encrypt. %v", tp1)
	return
}

// EncryptAfterCounter returns a count of finished RingMock.Encrypt invocations
func (mmEncrypt *RingMock) EncryptAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmEncrypt.afterEncryptCounter)
}

// EncryptBeforeCounter returns a count of RingMock.Encrypt invocations
func (mmEncrypt *RingMock) EncryptBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmEncrypt.beforeEncryptCounter)
}

// Calls returns a list of arguments used in each call to RingMock.Encrypt.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmEncrypt *mRingMockEncrypt) Calls() []*RingMockEncryptParams {
	mmEncrypt.mutex.RLock()

	argCopy := make([]*RingMockEncryptParams, len(mmEncrypt.callArgs))
	copy(argCopy, mmEncrypt.callArgs)

	mmEncrypt.mutex.RUnlock()

	return argCopy
}

// MinimockEncryptDone returns true if the count of the Encrypt invocations corresponds
// the number of defined expectations
func (m *RingMock) MinimockEncryptDone() bool {
	for _, e := range m.EncryptMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.EncryptMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterEncryptCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcEncrypt != nil && mm_atomic.LoadUint64(&m.afterEncryptCounter) < 1 {
		return false
	}
	return true
}

// MinimockEncryptInspect logs each unmet expectation
func (m *RingMock) MinimockEncryptInspect() {
	for _, e := range m.EncryptMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RingMock.Encrypt with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.EncryptMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterEncryptCounter) < 1 {
		if m.EncryptMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RingMock.Encrypt")
		} else {
			m.t.Errorf("Expected call to RingMock.Encrypt with params: %#v", *m.EncryptMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcEncrypt != nil && mm_atomic.LoadUint64(&m.afterEncryptCounter) < 1 {
		m.t.Error("Expected call to RingMock.Encrypt")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RingMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDecryptInspect()

		m.MinimockEncryptInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RingMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *RingMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDecryptDone() &&
		m.MinimockEncryptDone()
}
