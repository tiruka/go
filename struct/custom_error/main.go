package main

import "fmt"

type UserNotFound struct {
	UserName string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not Found %v\n", e.UserName)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{UserName: "mike"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
