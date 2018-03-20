package Error

import (
	"bytes"
	"fmt"
)
import "container/list"

type Result interface {
	GetErrors() *list.List
	IsSuccess() bool
	ErrorInfo() string
}

type Failure struct {
	info *list.List
}

type Success struct {
}

func NewFailure(s string) Failure {
	msg := list.New()
	msg.PushBack(s)
	return Failure{msg}
}

func (r Failure) ErrorInfo() string {
	var buffer bytes.Buffer
	for e := r.info.Front(); e != nil; e = e.Next() {
		buffer.WriteString(fmt.Sprintf("%v;", e.Value))
	}
	return buffer.String()
}

func (r Success) ErrorInfo() string {

	return ""
}

func (r Failure) IsSuccess() bool {
	return false
}

func (r Success) IsSuccess() bool {
	return true
}

func (r Failure) GetErrors() *list.List {
	return r.info
}

func (r Success) GetErrors() *list.List {
	return list.New()
}

func AccumulateResult(results []Result) Result {
	s := list.New()
	for _, result := range results {
		s.PushBackList(result.GetErrors())
	}
	if s.Len() == 0 {
		fmt.Printf("return Success\n")
		return Success{}
	} else {
		fmt.Printf("return Failure\n")
		return Failure{s}
	}
}
