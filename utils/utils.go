package utils

import "fmt"

// CheckError is a function to check for errors
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
