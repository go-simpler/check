// Package check provides convenience helpers to perform validations of any kind.
//
// Use That/Thatf to write conditions to check, multiple calls can be chained.
// The last call in the chain must be either FirstError or AllErrors.
package check

import "fmt"

// That checks whether the condition is true, and if not, records the error.
func That(cond bool, err error) *State {
	return new(State).That(cond, err)
}

// Thatf checks whether the condition is true, and if not, creates an error from format and args, then records it.
func Thatf(cond bool, format string, args ...any) *State {
	return new(State).Thatf(cond, format, args...)
}

// State holds the errors of the failed conditions.
// It is exported only for the purpose of documentation.
type State struct {
	errs []error
}

// That checks whether the condition is true, and if not, records the error.
func (s *State) That(cond bool, err error) *State {
	if err == nil {
		panic("check: a nil error provided")
	}
	if !cond {
		s.errs = append(s.errs, err)
	}
	return s
}

// Thatf checks whether the condition is true, and if not, creates an error from format and args, then records it.
func (s *State) Thatf(cond bool, format string, args ...any) *State {
	return s.That(cond, fmt.Errorf(format, args...))
}

// FirstError returns the error of the first failed condition.
func (s *State) FirstError() error {
	if len(s.errs) > 0 {
		return s.errs[0]
	}
	return nil
}

// AllErrors returns the errors of all failed conditions.
func (s *State) AllErrors() []error { return s.errs }
