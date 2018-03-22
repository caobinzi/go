package validation

type Success struct{}

func (r *Success) ErrorInfo() string {
	return ""
}

func (r *Success) IsOK() bool {
	return true
}
func (r *Success) Errors() []string {
	return []string{}
}
func NewSuccess() *Success {
	return &Success{}
}
