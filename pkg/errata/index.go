package errata

import (
	"fmt"
)

type ErrataParams struct {
	ErrMsg string
	Err    error
}

func CreateError(errMessage string, err error) error {
	return fmt.Errorf(errMessage, "caused error:", err)
}
