package main

import "validation"
import "fmt"

func checkUser(user string) validation.Result {
	return validation.NewFailure(fmt.Sprintf("Invalid User:%s", user))
}

func main() {
	var results validation.Results = []validation.Result{checkUser("abc"), checkUser("def"), checkUser("def")}
	result := results.Sum()
	//	validation.checkUser("s")
	if result.IsOK() {
		fmt.Printf("All good to go\n")
	} else {
		fmt.Printf("Erros found: %s\n", result.ErrorInfo())
	}
}
