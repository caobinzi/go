package main

import (
	"bytes"
	"fmt"
)
import "container/list"

type Result interface {
	getErrors() *list.List
	isSuccess() bool
	errorInfo() string
}

type Failure struct {
	info *list.List
}
type Success struct {
}

func (r Failure) errorInfo() string {
	var buffer bytes.Buffer
	for e := r.info.Front(); e != nil; e = e.Next() {
		buffer.WriteString(fmt.Sprintf("%v;", e.Value))
	}
	return buffer.String()
}

func (r Success) errorInfo() string {

	return ""
}

func (r Failure) isSuccess() bool {
	return false
}

func (r Success) isSuccess() bool {
	return true
}

func (r Failure) getErrors() *list.List {
	return r.info
}

func (r Success) getErrors() *list.List {
	return list.New()
}

func checkUser(user string) Result {
	s := list.New()
	s.PushBack(fmt.Sprintf("Invalid User:%s", user))
	return Failure{s}
}

func accumulateResult(results []Result) Result {
	s := list.New()
	for _, result := range results {
		s.PushBackList(result.getErrors())
	}
	if s.Len() == 0 {
		fmt.Printf("return Success\n")
		return Success{}
	} else {
		fmt.Printf("return Failure\n")
		return Failure{s}
	}
}

func main() {
	result := accumulateResult([]Result{checkUser("abc"), checkUser("def"), checkUser("def")})
	fmt.Printf("Result = %t\n", result.isSuccess())
	fmt.Printf("Error = %s\n", result.errorInfo())
}
