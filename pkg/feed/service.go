package feed

// Service provides basic operations on for feed service.
type Service interface {
}

type service struct {
	feed Repository
}

// NewService creates an instance of the service for feed service with necessary dependencies.
func NewService(feed Repository) Service {
	return &service{
		feed: feed,
	}
}
