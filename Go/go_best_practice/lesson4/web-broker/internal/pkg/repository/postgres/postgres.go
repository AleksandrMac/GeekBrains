package postgres

import (
	"github.com/vlslav/web-broker/internal/pkg/repository/postgres/alex"
)

type PgRepo struct {
	alex.Req
}

func New() *PgRepo {
	return &PgRepo{}
}
