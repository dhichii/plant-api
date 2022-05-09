package auth

// Outgoing port for auth
type Service interface {
	Login(email, password string) (string, error)
}
