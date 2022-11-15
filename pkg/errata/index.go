package errata

import (
	"fmt"
	"strings"
)

type ErrMessage struct {
	Text string
}

func CreateError(err error, errText []string) error {
	return fmt.Errorf(strings.Join(errText, "\n"), "caused error:", err)
}
