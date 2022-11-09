package errata

import (
	"fmt"
)

type ErrMessage struct {
	Text string
}

func CreateError(err error, errMessage ErrMessage) error {
	return fmt.Errorf(errMessage.Text, "caused error:", err)
}
