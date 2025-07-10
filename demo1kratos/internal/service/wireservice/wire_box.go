package wireservice

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/orzkratos/demokratos/demo1kratos/internal/conf"
	"github.com/orzkratos/demokratos/demo1kratos/internal/service"
)

type WireBox struct {
	GreeterService *service.GreeterService
}

func newWireBox(greeterService *service.GreeterService) *WireBox {
	return &WireBox{
		GreeterService: greeterService,
	}
}

func NewWireBox(confData *conf.Data, logger log.Logger) (*WireBox, func(), error) {
	return wireApp(confData, logger)
}
