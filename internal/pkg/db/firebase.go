package db

import (
	"context"
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
	return nil, nil
}
