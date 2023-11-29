package domain

type Comment struct {
	Id        uint
	Message   string
	CreatedAt string
	Creator   Creator
}
