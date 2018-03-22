package validation

import "strings"

type Failure struct {
	errors []string
}

func NewFailure(failMsg string) *Failure {
	return &Failure{
		errors: []string{failMsg},
	}
}

func (r *Failure) ErrorInfo() string {
	return strings.Join(r.errors, ",")
}

func (r *Failure) IsOK() bool {
	return false
}

func (r *Failure) Errors() []string {
	return r.errors
}
