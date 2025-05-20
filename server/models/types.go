package models

type Model interface {
	GetID() string
	SetID(string)
}
