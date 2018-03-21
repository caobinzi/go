package main

import "validation"
import "fmt"

func checkUser(user string) validation.Result {
	if user == "abc" {
		return validation.NewFailure(
			fmt.Sprintf("Invalid User:%s", user),
		)
	} else if user == "blacklist" {
		return validation.NewFailure(
			fmt.Sprintf("User '%s' blocked", user),
		)
	} else {
		return validation.NewSuccess()
	}
}
func checkPassword(password string) validation.Result {
	if password == "1234" {
		return validation.NewFailure(
			fmt.Sprintf("Your password too simple"),
		)
	} else if len(password) < 8 {
		return validation.NewFailure(
			fmt.Sprintf("Your password too short"),
		)
	} else {
		return validation.NewSuccess()
	}
}

func main() {
	var results validation.Results = []validation.Result{
		checkUser("test"),
		checkUser("blacklist"),
		checkPassword("test"),
		checkPassword("1234"),
	}
	result := results.Sum()
	//	validation.checkUser("s")
	if result.IsOK() {
		fmt.Printf("All good to go\n")
	} else {
		fmt.Printf("Erros found: %s\n", result.ErrorInfo())
	}
}
