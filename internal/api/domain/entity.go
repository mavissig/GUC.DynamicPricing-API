package domain

import "github.com/google/uuid"

type Data struct {
	Id   uuid.UUID `json:"id"`
	Data []byte    `json:"data"`
}
