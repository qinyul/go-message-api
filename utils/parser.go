package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func UUIDParser(val string) uuid.UUID {
	parsedUUID, err := uuid.Parse(val)
	if err != nil {
		fmt.Printf("Failed to parsing uuid: %v\n", err)
		return uuid.Nil
	}
	return parsedUUID
}
