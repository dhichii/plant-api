package native

// Ingoing port for native
type Repository interface {
	Create(*Native) error
	GetAll() ([]Native, error)
	Get(id int) (*Native, error)
	GetByName(string) (*Native, error)
}

// Outgoing port for native
type Service interface {
	Create(*Native) error
	GetAll() ([]Native, error)
	Get(id int) (*Native, error)
	GetByName(name string) (*Native, error)
}