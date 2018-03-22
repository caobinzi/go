package validation

import "github.com/cheekybits/genny/generic"

type Generic generic.Type

type GenericSuccess struct {
	valueGeneric Generic
}

func (r *GenericSuccess) ErrorInfo() string {
	return ""
}

func (r *GenericSuccess) IsOK() bool {
	return true
}
func (r *GenericSuccess) Errors() []string {
	return []string{}
}
func NewGenericSuccess(a Generic) *GenericSuccess {
	return &GenericSuccess{valueGeneric: a}
}
func (r *GenericSuccess) GetGeneric() Generic {
	return r.valueGeneric
}
