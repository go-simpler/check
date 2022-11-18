# check

[![ci](https://github.com/go-simpler/check/actions/workflows/go.yml/badge.svg)](https://github.com/go-simpler/check/actions/workflows/go.yml)
[![docs](https://pkg.go.dev/badge/github.com/go-simpler/check.svg)](https://pkg.go.dev/github.com/go-simpler/check)
[![report](https://goreportcard.com/badge/github.com/go-simpler/check)](https://goreportcard.com/report/github.com/go-simpler/check)
[![codecov](https://codecov.io/gh/go-simpler/check/branch/main/graph/badge.svg)](https://codecov.io/gh/go-simpler/check)

Convenience helpers to perform validations of any kind

## 📦 Install

```shell
go get github.com/go-simpler/check
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
