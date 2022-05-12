package plantnative

// Ingoing port for plant native
type Repository interface {
	Delete(id int) error
}

// Outgoing port for plant native
type Service interface {
	Delete(id int) error
}
