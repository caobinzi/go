// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package data

type UserSuccess struct {
	valueUser User
}

func (r *UserSuccess) ErrorInfo() string {
	return ""
}

func (r *UserSuccess) IsOK() bool {
	return true
}
func (r *UserSuccess) Errors() []string {
	return []string{}
}
func NewUserSuccess(a User) *UserSuccess {
	return &UserSuccess{a}
}
func (r *UserSuccess) GetUser() User {
	return r.valueUser
}
