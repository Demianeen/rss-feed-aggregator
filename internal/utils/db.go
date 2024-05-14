package utils

import (
	"time"

	"github.com/google/uuid"
)

// automates the creation of standard fields for database entries
func NewDbEntry() (id uuid.UUID, created_at time.Time, updated_at time.Time) {
	now := time.Now().UTC()
	return uuid.New(), now, now
}
