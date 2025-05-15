package main

import (
	"fmt"

	"github.com/atomhudson/go_lang/auth"
	"github.com/atomhudson/go_lang/user"
	"github.com/fatih/color"
)

// In this way : go get github.com/fatih/color
// We can install the package using go get command
func main() {
	auth.LoginWithCredentials("admin", "password")
	session := auth.GetSession()
	fmt.Println("Session:", session)

	fmt.Println("--------------------------------------------------------------")
	user := user.User{
		ID:       1,
		Username: "admin",
		Password: "password",
		Email:    "admin@gmail.com",
	}
	color.Cyan("User ID:", user.ID)
	fmt.Println()
	color.HiGreen("Username:", user.Username)
	fmt.Println()
	color.Red("Password:", user.Password)
	fmt.Println()
	color.Magenta("Email:", user.Email)
	fmt.Println()
	color.Yellow("--------------------------------------------------------------")
	fmt.Println()
	auth.LoginWithCredentials(user.Username, user.Password)
}
