package auth

import "fmt"

func LoginWithCredentials(username string, password string) (bool, error) {
	if username == "admin" && password == "password" {
		fmt.Println("Login successful!")
		return true, nil
	}
	fmt.Println("Login failed!")
	return false, nil
}
