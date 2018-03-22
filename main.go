package main

import (
	"fmt"
	"github.com/caobinzi/go/data"
	"github.com/caobinzi/go/validation"
	"reflect"
)

func sum(s validation.Results, c chan validation.Result) {
	for _, v := range s {
		c <- v
	}
}

func main() {
	var results validation.Results = []validation.Result{
		data.CheckUser("binzi"),
		data.CheckPermission("binzi"),
		data.CheckPassword("ppppppppp"),
		data.GetUserID("binzi"),
		data.GetUser("binzi"),
	}
	result := results.Sum()
	c := make(chan validation.Result)
	go sum(results, c)
	x, y := <-c, <-c
	fmt.Printf("First 2 results:%s -> %s\n", reflect.ValueOf(x).Kind(), reflect.ValueOf(y).Kind())
	//	validation.checkUser("s")
	if result.IsOK() {
		fmt.Printf("All good to go\n")
	} else {
		fmt.Printf("Erros found: %s\n", result.ErrorInfo())
	}
}
