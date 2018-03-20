package main

import "Error"
import "fmt"

func checkUser(user string) Error.Result {
	return Error.NewFailure(fmt.Sprintf("Invalid User:%s", user))
}
func main() {
	result := Error.AccumulateResult([]Error.Result{checkUser("abc"), checkUser("def"), checkUser("def")})
	//	Error.checkUser("s")
	fmt.Printf("Result = %t\n", result.IsSuccess())
	fmt.Printf("Error = %s\n", result.ErrorInfo())
}
