package utils

import "fmt"

func ReturnError(errFunc string, err error) error {
	fmt.Println("\n********** uh oh very virus")
	fmt.Println("\n", "error while calling ", errFunc, "\n", err)

	return err
}
