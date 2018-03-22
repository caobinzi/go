package main

import (
	"fmt"

	"github.com/caobinzi/go/data"
	"github.com/caobinzi/go/validation"
)

func main() {
	var results validation.Results = []validation.Result{
		data.CheckUser("binzi"),
		data.CheckPermission("binzi"),
		data.CheckPassword("ppppppppp"),
		data.GetUserID("binzi"),
	}
	result := results.Sum()
	//	validation.checkUser("s")
	if result.IsOK() {
		fmt.Printf("All good to go\n")
	} else {
		fmt.Printf("Erros found: %s\n", result.ErrorInfo())
	}
}
