package data

import "github.com/cheekybits/genny/generic"

type Something generic.Type

type SomethingSuccess struct {
	value Something
}

func (r *SomethingSuccess) ErrorInfo() string {
	return ""
}

func (r *SomethingSuccess) IsOK() bool {
	return true
}
func (r *SomethingSuccess) Errors() []string {
	return []string{}
}
func NewStringSuccess(a Something) *SomethingSuccess {
	return &SomethingSuccess{value: a}
}

func (r *SomethingSuccess) GetInt() int {
	return r.value.(int)
}
