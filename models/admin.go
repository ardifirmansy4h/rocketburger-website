package models

type Admin struct {
	ID int
	Nama string
	Email string
	Password string
}

type InputAdmin struct {
	Nama string
	Email string
	Password string
}