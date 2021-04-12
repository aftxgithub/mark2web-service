package db

import (
	"bytes"
	"testing"
)

func testDBImplementation(t *testing.T, dbImpl DB) {
	testURL := "http://mark2web.test/3598350bd2279289b076e3faecec71eadad4d17e"
	testHTML := []byte("<h1>Hello World</h1>")

	dbImpl.Save(testURL, testHTML)

	gotHTML := dbImpl.GetHTMLFor(testURL)
	if !bytes.Equal(testHTML, gotHTML) {
		t.Errorf("expected retrieved content to be %s, got %s", testHTML, gotHTML)
	}
}
