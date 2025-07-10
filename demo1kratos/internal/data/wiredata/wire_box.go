package wiredata

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
)

type WireBox struct {
	GreeterRepo biz.GreeterRepo
}

func newWireBox(greeterRepo biz.GreeterRepo) *WireBox {
	return &WireBox{
		GreeterRepo: greeterRepo,
	}
}

func NewWireBox(confData *conf.Data, logger log.Logger) (*WireBox, func(), error) {
	return wireApp(confData, logger)
}
