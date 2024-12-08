package user

type Repository interface {
	GetByEmail(email string) (*User, error)
	GetById(id string) (*User, error)
}