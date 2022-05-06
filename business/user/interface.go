package user

// Ingoing port for user
type Repository interface {
	Create(user User) error
	GetAll() ([]User, error)
	Get(id int) (*User, error)
	Update(id int, user User) error
}

// Outgoing port for user
type Service interface {
	Create(user User) error
	GetAll() ([]User, error)
	Get(id int) (*User, error)
	Update(id int, user User) error
}