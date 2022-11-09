package appUtils

import "github.com/twinj/uuid"

func CreateUuid() string {
	return uuid.NewV4().String()
}
