package data

import (
	"fmt"

	"github.com/caobinzi/go/validation"
)

func CheckUser(user string) validation.Result {
	if user == "abc" {
		return validation.NewFailure(
			fmt.Sprintf("Invalid User:%s", user),
		)
	} else if user == "blacklist" {
		return validation.NewFailure(
			fmt.Sprintf("User '%s' blocked", user),
		)
	} else {
		return validation.NewSuccess(nil)
	}
}
func CheckPermission(user string) validation.Result {
	if user == "abc" {
		return validation.NewFailure(
			fmt.Sprintf("Read only user found :%s", user),
		)
	} else if user == "root" {
		return validation.NewFailure(
			fmt.Sprintf("Root user '%s' is not allowed", user),
		)
	} else {
		return validation.NewSuccess(nil)
	}
}

func CheckPassword(password string) validation.Result {
	if password == "1234" {
		return validation.NewFailure(
			fmt.Sprintf("Your password too simple"),
		)
	} else if len(password) < 8 {
		return validation.NewFailure(
			fmt.Sprintf("Your password too short"),
		)
	} else {
		return validation.NewSuccess(nil)
	}
}

func GetUserID(user string) validation.Result {
	if user == "binzi" {
		return validation.NewSuccess(0)
	} else {
		return validation.NewFailure(
			fmt.Sprintf("User does not exit"),
		)
	}
}
