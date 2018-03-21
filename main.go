package main

import "validation"
import "fmt"

func checkUser(user string) validation.Result {
	if user == "abc" {
		return validation.NewFailure(
			fmt.Sprintf("Invalid User:%s", user),
		)
	} else if user == "def" {
		return validation.NewFailure(
			fmt.Sprintf("Invalid User:%s", user),
		)
	} else {
		return validation.NewSuccess()
	}
}

func main() {
	var results validation.Results = []validation.Result{
		checkUser("test"),
		checkUser("test"),
		checkUser("test"),
	}
	result := results.Sum()
	//	validation.checkUser("s")
	if result.IsOK() {
		fmt.Printf("All good to go\n")
	} else {
		fmt.Printf("Erros found: %s\n", result.ErrorInfo())
	}
}
