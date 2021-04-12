package db

import (
	"bytes"
	"testing"
)

func testDBImplementation(t *testing.T, dbImpl DB) {
	testID := "3598350bd2279289b076e3faecec71eadad4d17e"
	testHTML := []byte("<h1>Hello World</h1>")

	err := dbImpl.Save(testID, testHTML)
	if err != nil {
		t.Error(err)
	}

	gotHTML, err := dbImpl.GetHTMLFor(testID)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(testHTML, gotHTML) {
		t.Errorf("expected retrieved content to be '%s', got '%s'", testHTML, gotHTML)
	}
}
