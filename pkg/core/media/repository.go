package media

// Repository provides access to the media store.
type Repository interface {
	Create(media *Media) error
}
