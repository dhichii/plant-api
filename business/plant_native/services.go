package plantnative

type service struct {
	repo Repository
}

// Construct plant native service object
func NewService(repo Repository) Service {
	return &service{repo}
}

// Delete plant native by given id
func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}