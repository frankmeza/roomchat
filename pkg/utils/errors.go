package utils

import "fmt"

func ReturnError(functionInError string, err error) error {
	fmt.Println("\n********** uh oh very virus")
	fmt.Println("\n", "error while calling ", functionInError, "\n", err)

	return err
}
