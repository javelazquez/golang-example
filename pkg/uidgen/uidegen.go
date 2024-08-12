package uidgen

import "github.com/google/uuid"

type UIDGen interface {
	NewID() string
}

type uidgen struct{}

func New() UIDGen {
	return &uidgen{}
}

func (u uidgen) NewID() string {
	return uuid.New().String()
}
