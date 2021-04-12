package db

// DB defines the API for any database implementation to be used in the mark2web service
type DB interface {
	Save(URL string, HTML []byte) error
	GetHTMLFor(URL string) []byte
}
