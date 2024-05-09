package firebases

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"os"
)

type (
	Firebase interface {
		Store(ctx context.Context, path string) error
	}

	FirebaseConf struct {
		app *firebase.App
	}
)

func NewFirebaseConf(ctx context.Context) (*FirebaseConf, error) {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &FirebaseConf{
		app: app,
	}, nil
}

func (f *FirebaseConf) Store(ctx context.Context, path string) error {
	store, err := f.app.Storage(ctx)
	if err != nil {
		return err
	}
	bucket, err := store.Bucket("gs://bulletinboard-37277.appspot.com")
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	wc := bucket.Object(path).NewWriter(ctx)
	if _, err := wc.Write([]byte{}); err != nil {
		return err
	}
	return nil
}
