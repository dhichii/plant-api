package account

type service struct {
	repository Repository
}

// Counstruct account service object
func NewService(repo Repository) Service {
	return &service{repo}
}

// Update user email in database
func (s *service) UpdateEmail(id int, email string) error {
	result, err := s.repository.GetEmailByID(id)
	if err != nil {
		return err
	}

	if result == email {
		return nil
	}

	return s.repository.UpdateEmail(id, email)
}

// Update user password in database
func (s *service) UpdatePassword(id int, password string) error {
	return s.repository.UpdatePassword(id, password)
}
