package entity

import "github.com/google/uuid"

type Media interface {
	GetID() uuid.UUID
	GetName() string
}
