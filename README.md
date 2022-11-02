# check

Convenience helpers to perform validations of any kind

## 📦 Install

```shell
go get github.com/junk1tm/check
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
	
for _, err := range errs {
	// handle error
}
```
