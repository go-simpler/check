// Package check implements functions for writing validations declaratively.
//
// Use [That]/[Thatf] to write conditions to check, multiple calls can be chained.
// The last call in the chain must be FirstError, AllErrors, or JoinErrors.
package check

import (
	"errors"
	"fmt"
)

// That checks whether the condition is true, and if not, records the error.
func That(cond bool, err error) *State {
	return new(State).That(cond, err)
}

// Thatf checks whether the condition is true, and if not, creates an error from format and args, then records it.
func Thatf(cond bool, format string, args ...any) *State {
	return new(State).Thatf(cond, format, args...)
}

// State holds the recorded errors.
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

// FirstError returns the first recorded error.
func (s *State) FirstError() error {
	if len(s.errs) > 0 {
		return s.errs[0]
	}
	return nil
}

// AllErrors returns all the recorded errors.
func (s *State) AllErrors() []error { return s.errs }

// JoinErrors returns all the recorded errors joined via [errors.Join].
func (s *State) JoinErrors() error { return errors.Join(s.errs...) }
