package check_test

import (
	"errors"
	"testing"

	"github.com/junk1tm/check"
)

func TestCheck(t *testing.T) {
	t.Run("first error", func(t *testing.T) {
		err12 := errors.New("1 and 2 are not equal")
		err34 := errors.New("3 and 4 are not equal")

		err := check.
			That(1 == 2, err12).
			That(3 == 4, err34).
			FirstError()

		if !errors.Is(err, err12) {
			t.Errorf("got %v; want %v", err, err12)
		}
	})

	t.Run("all errors", func(t *testing.T) {
		errs := check.
			Thatf("foo" == "baz", "foo and bar are not equal").
			Thatf(true == false, "true and false are not equal").
			AllErrors()

		if len(errs) != 2 {
			t.Fatalf("want 2 errors")
		}

		if errs[0] == nil || errs[1] == nil {
			t.Errorf("want all errors to be not nil")
		}
	})

	// TODO(jun1ktm): test That's panic case
	// TODO(jun1ktm): test FirstError's nil case
}