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
		data.GetUser("binzi"),
	}
	result := results.Sum()

	if result.IsOK() {
		fmt.Printf("All good to go:%s\n", data.GetUser("binzi").(*data.UserSuccess).GetUser().Name)
	} else {
		fmt.Printf("Erros found: %s\n", result.ErrorInfo())
	}
}
