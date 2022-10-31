package errata

import (
	"errors"
	"fmt"
)

type ErrataParams struct {
	ErrMsg string
	Err    error
}

func CreateError(params ErrataParams) error {
	errorMessage := fmt.Sprintf(params.ErrMsg, "caused error:", params.Err)
	return errors.New(errorMessage)
}
