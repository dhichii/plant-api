package plantnative

// Ingoing port for plant native
type Repository interface {
	Delete(plantID int) error
}

// Outgoing port for plant native
type Service interface {
	Delete(plantID int) error
}
