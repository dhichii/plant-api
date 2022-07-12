package account

// Ingoing port for user
type Repository interface {
	GetEmailByID(id int) (string, error)
	UpdateEmail(id int, email string) error
	UpdatePassword(id int, password string) error
}

// Ongoing port for account
type Service interface {
	UpdateEmail(id int, email string) error
	UpdatePassword(id int, password string) error
}
