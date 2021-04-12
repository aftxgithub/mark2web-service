package db

import "testing"

func TestFSDBImplementation(t *testing.T) {
	testDBImplementation(t, new(FSDatabase))
}
