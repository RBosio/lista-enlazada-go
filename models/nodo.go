package models

type Nodo struct {
	User *User
	Next *Nodo
}
