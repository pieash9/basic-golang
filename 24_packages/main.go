package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/pieash9/basic-golang/24_packages/auth"
	"github.com/pieash9/basic-golang/24_packages/user"
)

func main() {
	auth.LoginWithCredentials("pieash", "123")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "pieash@gmail.com",
		Name:  "pieash",
	}

	// fmt.Println(user.Email, user.Name)
	color.Green(user.Email)
}
