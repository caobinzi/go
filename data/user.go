package data

import (
	"fmt"

	"github.com/caobinzi/go/validation"
)

type User struct {
	name string
	age  int
}

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
		return validation.NewGenericSuccess(nil)
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
		return validation.NewGenericSuccess(nil)
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
		return validation.NewGenericSuccess(nil)
	}
}

func GetUserID(user string) validation.Result {
	if user == "binzi" {
		return validation.NewGenericSuccess(0)
	} else {
		return validation.NewFailure(
			fmt.Sprintf("User does not exit"),
		)
	}
}
func GetUser(id string) validation.Result {
	if id == "binzi" {
		return NewUserSuccess(
			User{
				name: "binzi",
				age:  20,
			},
		)
	} else {
		return validation.NewFailure(
			fmt.Sprintf("User does not exit"),
		)
	}
}
