package wirebiz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
)

type WireBox struct {
	GreeterUsecase *biz.GreeterUsecase
}

func newWireBox(greeterUsecase *biz.GreeterUsecase) *WireBox {
	return &WireBox{
		GreeterUsecase: greeterUsecase,
	}
}

func NewWireBox(confData *conf.Data, logger log.Logger) (*WireBox, func(), error) {
	return wireApp(confData, logger)
}
