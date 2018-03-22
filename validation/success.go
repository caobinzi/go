package validation

type Any interface{}

type Success struct {
	value Any
}

func (r *Success) ErrorInfo() string {
	return ""
}

func (r *Success) IsOK() bool {
	return true
}
func (r *Success) Errors() []string {
	return []string{}
}
func NewSuccess(a Any) *Success {
	return &Success{value: a}
}

func (r *Success) GetInt() int {
	return r.value.(int)
}
