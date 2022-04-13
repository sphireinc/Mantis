package uuid

import (
	guid "github.com/google/uuid"
)

type uuid struct {
	UUID guid.UUID `json:"uuid"`
}

func New() uuid {
	return uuid{
		UUID: guid.New(),
	}
}

func (u uuid) String() string {
	return u.UUID.String()
}
