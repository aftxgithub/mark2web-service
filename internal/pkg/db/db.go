package db

// DB defines the API for any database implementation to be used in the mark2web service
type DB interface {
	Save(ID string, HTML []byte) error
	GetHTMLFor(ID string) ([]byte, error)
}
