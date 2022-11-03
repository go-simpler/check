// Package check provides convenience helpers to perform validations of any
// kind.
//
// Use That/Thatf to write conditions to check, multiple calls can be chained.
// The last call in the chain must be either FirstError or AllErrors.
package check

import "fmt"

// That checks whether cond is true, and if not, records the error. That panics
// if the error is nil.
func That(cond bool, err error) *state {
	return new(state).That(cond, err)
}

// Thatf checks whether cond is true, and if not, creates an error from format
// and args, then records it.
func Thatf(cond bool, format string, args ...any) *state {
	return new(state).Thatf(cond, format, args...)
}

// state holds the errors of the failed conditions.
type state struct {
	errs []error
}

// That checks whether cond is true, and if not, records the error. That panics
// if the error is nil.
func (s *state) That(cond bool, err error) *state {
	if err == nil {
		panic("check: a nil error is provided")
	}
	if !cond {
		s.errs = append(s.errs, err)
	}
	return s
}

// Thatf checks whether cond is true, and if not, creates an error from format
// and args, then records it.
func (s *state) Thatf(cond bool, format string, args ...any) *state {
	return s.That(cond, fmt.Errorf(format, args...))
}

// FirstError returns the error of the first failed condition.
func (s *state) FirstError() error {
	if len(s.errs) > 0 {
		return s.errs[0]
	}
	return nil
}

// AllErrors returns the errors of all failed conditions.
func (s *state) AllErrors() []error { return s.errs }
