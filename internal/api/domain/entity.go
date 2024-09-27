package domain

import "github.com/google/uuid"

type Data struct {
	ID   uuid.UUID `json:"id"`
	Data []byte    `json:"data"`
}
