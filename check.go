// Package check provides convenience helpers to perform validations of any
// kind.
//
// Use That/Thatf to write conditions to check, multiple calls can be chained.
// The last call in the chain must be either FirstError or AllErrors.
package check

import "fmt"

// That checks whether cond is true, and if not, records the error. That panics
// if the error is nil.
func That(cond bool, err error) *check {
	return new(check).That(cond, err)
}

// Thatf checks whether cond is true, and if not, creates an error from format
// and args, then records it.
func Thatf(cond bool, format string, args ...any) *check {
	return That(cond, fmt.Errorf(format, args...))
}

// check holds the conditions to check and their corresponding errors.
type check struct {
	conds []bool
	errs  []error
}

// That checks whether cond is true, and if not, records the error. That panics
// if the error is nil.
func (ch *check) That(cond bool, err error) *check {
	if err == nil {
		panic("check: a nil error is provided")
	}
	ch.conds = append(ch.conds, cond)
	ch.errs = append(ch.errs, err)
	return ch
}

// Thatf checks whether cond is true, and if not, creates an error from format
// and args, then records it.
func (ch *check) Thatf(cond bool, format string, args ...any) *check {
	return ch.That(cond, fmt.Errorf(format, args...))
}

// FirstError returns the error of the first failed condition.
func (ch *check) FirstError() error {
	for i := range ch.conds {
		if !ch.conds[i] {
			return ch.errs[i]
		}
	}
	return nil
}

// AllErrors returns the errors of all failed conditions.
func (ch *check) AllErrors() []error {
	var errs []error
	for i := range ch.conds {
		if !ch.conds[i] {
			errs = append(errs, ch.errs[i])
		}
	}
	return errs
}
