package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

func GetNewUUID() string {
	// or error handling
	u2, err := uuid.NewUUID()
	if err != nil {
		// Fallback to random integer. Not expected as V4 uuid is based on random number
		randNum := rand.Int()
		return fmt.Sprintf("%d", randNum)
	}
	return fmt.Sprintf("%s", u2)
}