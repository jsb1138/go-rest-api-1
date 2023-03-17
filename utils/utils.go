package utils

import "fmt"

// CheckError is a function to check for errors
// Eventually I stopped using it because it was too generic
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
