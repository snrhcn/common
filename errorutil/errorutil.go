package errorutil

import "strings"

/*
CompositeError can collect multiple errors in a single error object.
*/
type CompositeError struct {
	Errors []string
}

/*
NewCompositeError creates a new composite error object.
*/
func NewCompositeError() *CompositeError {
	return &CompositeError{make([]string, 0)}
}

/*
Add adds an error.
*/
func (ce *CompositeError) Add(e error) {
	ce.Errors = append(ce.Errors, e.Error())
}

/*
HasErrors returns true if any error have been collected.
*/
func (ce *CompositeError) HasErrors() bool {
	return len(ce.Errors) > 0
}

/*
Error returns all collected errors as a string.
*/
func (ce *CompositeError) Error() string {
	return strings.Join(ce.Errors, "; ")
}
