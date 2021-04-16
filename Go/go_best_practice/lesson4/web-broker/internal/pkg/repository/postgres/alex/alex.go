package alex

import "github.com/vlslav/web-broker/internal/pkg/model"

type Req struct{}

func New() *Req {
	return &Req{}
}

func (*Req) Get(getReq string) (string, error) {
	return "", nil
}

func (*Req) Put(putReq *model.PutValue) error {
	return nil
}
