package gaea

// Service ...
type Service struct {
	Repo Repository
}

// NewService return a new Gaea service
func NewService(repo Repository) (*Service, error) {
	return &Service{
		Repo: repo,
	}, nil
}
