package postgres

import "github.com/vlslav/web-broker/internal/pkg/model"

type PgRepo struct{}

func New() *PgRepo {
	return &PgRepo{}
}

func (pgr *PgRepo) Get(getReq string) (string, error) {
	// TODO: impl
	return "", nil
}

func (pgr *PgRepo) Put(putReq *model.PutValue) error {
	// TODO: impl
	return nil
}
