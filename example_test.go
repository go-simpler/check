package check_test

import (
	"errors"
	"fmt"

	"github.com/junk1tm/check"
)

var user = struct {
	Name  string
	Age   int
	Email string
}{
	Name:  "",
	Age:   10,
	Email: "user@email",
}

func isEmail(string) bool { return false }

var errEmptyName = errors.New("name must not be empty")

func ExampleThat() {
	errs := check.
		That(user.Name != "", errEmptyName).
		Thatf(user.Age >= 18, "%d y.o. is too young", user.Age).
		Thatf(isEmail(user.Email), "%s is invalid email", user.Email).
		AllErrors() // OR FirstError() to check only the first error

	for _, err := range errs {
		fmt.Println(err)
	}

	// Output:
	// name must not be empty
	// 10 y.o. is too young
	// user@email is invalid email
}
