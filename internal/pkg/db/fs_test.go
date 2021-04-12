package db

import (
	"os"
	"testing"
)

func TestFSDBImplementation(t *testing.T) {
	testDBImplementation(t, &FSDatabase{os.TempDir()})
}
