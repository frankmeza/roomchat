package errata

import (
	"errors"
	"fmt"
)

type ErrataParams struct {
	ErrFunc string
	Err     error
}

func CreateError(params ErrataParams) error {
	errorMessage := fmt.Sprintf(params.ErrFunc, "caused error:", params.Err)
	return errors.New(errorMessage)
}
