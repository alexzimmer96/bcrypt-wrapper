[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT) 
[![codecov](https://codecov.io/gh/alexzimmer96/bcrypt-wrapper/branch/master/graph/badge.svg)](https://codecov.io/gh/alexzimmer96/bcrypt-wrapper) 
[![Go Report Card](https://goreportcard.com/badge/github.com/alexzimmer96/bcrypt-wrapper)](https://goreportcard.com/report/github.com/alexzimmer96/bcrypt-wrapper) 
[![GoDoc](https://godoc.org/github.com/alexzimmer96/bcrypt-wrapper?status.svg)](https://godoc.org/github.com/alexzimmer96/bcrypt-wrapper) 

# BCrypt-Wrapper
BCrypt-Wrapper is a small wrapper around Go's standard BCrypt-Implementation with the goal of increasing the cost-factor when its needed.
To do this, everytime you use this wrapper when you verify its password, there is a check if the used cost-factor is not already to low.
When the cost-factor is to low, the password will be hashed again with the focused cost-factor and returned.

## Example Usage
```go
wrapper := NewBCryptWrapper(14) // Create a new BCrypt-Wrapper with cost factor 14
someOutdatedHash := "$2y$11$KPR/tNhxP0RQcO7gSNjgJuXwu1jrwkfuEt2resN98faTtNnzq0DMa" // Hashed with cost-factor 14
someMatchingPassword := "E&dWBjxaE*8V"
newpass, err := wrapper.CompareHashAndPassword([]byte(someOutdatedHash), []byte(someMatchingPassword)) // Would return no error but the password hashed with cost-factor 14
```