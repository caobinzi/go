package validation

type Result interface {
	Errors() []string
	IsOK() bool
	ErrorInfo() string
}

type Results []Result

func (results Results) Sum() Result {
	errors := []string{}
	for _, result := range results {
		errors = append(errors, result.Errors()...)
	}
	if len(errors) == 0 {
		return &Success{}
	} else {
		return &Failure{errors}
	}
}
