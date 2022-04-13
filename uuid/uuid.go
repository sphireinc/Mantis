package uuid

import (
	guid "github.com/google/uuid"
)

// UUID our primary UUID struct
type UUID struct {
	UUID guid.UUID `json:"uuid"`
}

// New returns a newly created UUID
func New() UUID {
	return UUID{
		UUID: guid.New(),
	}
}

// String is a wrapper for Google's UUID.String
func (u UUID) String() string {
	return u.UUID.String()
}
