package uuid

import (
	guid "github.com/google/uuid"
)

// UUID our primary UUID struct
type UUID struct {
	UUID guid.UUID `json:"uuid"`
}

func New() UUID {
	return UUID{
		UUID: guid.New(),
	}
}

func (u UUID) String() string {
	return u.UUID.String()
}
