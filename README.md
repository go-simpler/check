# check

[![checks](https://github.com/go-simpler/check/actions/workflows/checks.yml/badge.svg)](https://github.com/go-simpler/check/actions/workflows/checks.yml)
[![pkg.go.dev](https://pkg.go.dev/badge/go-simpler.org/check.svg)](https://pkg.go.dev/go-simpler.org/check)
[![goreportcard](https://goreportcard.com/badge/go-simpler.org/check)](https://goreportcard.com/report/go-simpler.org/check)
[![codecov](https://codecov.io/gh/go-simpler/check/branch/main/graph/badge.svg)](https://codecov.io/gh/go-simpler/check)

Convenience helpers to perform validations of any kind

## 📦 Install

```shell
go get go-simpler.org/check
```

## 📋 Usage

Use `That`/`Thatf` to write conditions to check, multiple calls can be chained.
The last call in the chain must be either `FirstError` or `AllErrors`.

```go
errs := check.
	That(user.Name != "", errEmptyName).
	Thatf(user.Age >= 18, "%d y.o. is too young", user.Age).
	Thatf(isEmail(user.Email), "%s is invalid email", user.Email).
	AllErrors() // OR FirstError() to check only the first error
```

The same code without `check`:

```go
var errs []error
if user.Name == "" {
	errs = append(errs, errEmptyName)
}
if user.Age < 18 {
	errs = append(errs, fmt.Errorf("%d y.o. is too young", user.Age))
}
if !isEmail(user.Email) {
	errs = append(errs, fmt.Errorf("%s is invalid email", user.Email))
}
```
