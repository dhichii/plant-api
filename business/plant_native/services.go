package plantnative

type service struct {
	repo Repository
}

// Construct plant native service object
func NewService(repo Repository) Service {
	return &service{repo}
}

// Delete plant native by given plant id
func (s *service) Delete(plantID int) error {
	return s.repo.Delete(plantID)
}