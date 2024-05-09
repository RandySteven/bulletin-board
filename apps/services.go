package apps

import (
	"context"
	"task_mission/pkg/firebases"
)

type Services struct {
	Firebase firebases.Firebase
}

func NewServices(ctx context.Context) (*Services, error) {
	firebaseConf, err := firebases.NewFirebaseConf(ctx)
	if err != nil {
		return nil, err
	}
	return &Services{
		Firebase: firebaseConf,
	}, nil
}
