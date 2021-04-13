package db

import (
	"context"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
)

// FirebaseDB is a Firebase implementation of the DB interface
type FirebaseDB struct {
	bucket *storage.BucketHandle
	logger *log.Logger
}

func NewFirebaseDB(l *log.Logger) (*FirebaseDB, error) {
	log.Traceln("building new firebase db")

	config := &firebase.Config{
		StorageBucket: "mark2web.appspot.com",
	}
	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		return nil, err
	}
	client, err := app.Storage(context.Background())
	if err != nil {
		return nil, err
	}
	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, err
	}
	return &FirebaseDB{bucket, l}, nil
}

func (f *FirebaseDB) Save(ID string, HTML []byte) error {
	f.logger.Tracef("Saving %s at id '%s'", HTML, ID)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	wc := f.bucket.Object(ID).NewWriter(ctx)
	defer wc.Close()

	_, err := wc.Write(HTML)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirebaseDB) GetHTMLFor(ID string) ([]byte, error) {
	f.logger.Tracef("Getting HTML for id '%s'", ID)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	wc, err := f.bucket.Object(ID).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer wc.Close()

	HTMLBytes, err := ioutil.ReadAll(wc)
	if err != nil {
		return nil, err
	}

	return HTMLBytes, nil
}
