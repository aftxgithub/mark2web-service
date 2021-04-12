package db

import (
	"fmt"
	"io/ioutil"
)

// FSDatabase is a file-based implementation of the DB interface
type FSDatabase struct{}

func (f *FSDatabase) Save(ID string, HTML []byte) error {
	filename := fmt.Sprintf("%s.html", ID)
	err := ioutil.WriteFile(filename, HTML, 0664)
	return err
}

func (f *FSDatabase) GetHTMLFor(ID string) ([]byte, error) {
	filename := fmt.Sprintf("%s.html", ID)
	return ioutil.ReadFile(filename)
}
