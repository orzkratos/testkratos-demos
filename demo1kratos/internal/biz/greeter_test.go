package biz_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/orzkratos/demokratos/demo1kratos"
	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
	"github.com/orzkratos/demokratos/demo1kratos/internal/biz/wirebiz"
	"github.com/orzkratos/demokratos/demo1kratos/internal/pkg/cfgdata"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

var caseWireBox *wirebiz.WireBox

func TestMain(m *testing.M) {
	configPath := osmustexist.ROOT(filepath.Join(demo1kratos.SourceRoot(), "configs"))
	zaplog.LOG.Debug("run-test-main", zap.String("config-path", configPath))

	cfg := cfgdata.ParseConfig(configPath)
	wireBox, cleanup, err := wirebiz.NewWireBox(cfg.Data, log.NewStdLogger(os.Stdout))
	must.Done(err)
	defer cleanup()

	caseWireBox = wireBox
	zaplog.LOG.Debug("run-test-main")
	m.Run()
}

func TestGreeterUsecase_CreateGreeter(t *testing.T) {
	res, err := caseWireBox.GreeterUsecase.CreateGreeter(context.Background(), &biz.Greeter{
		Hello: "abc",
	})
	require.NoError(t, err)
	t.Log(neatjsons.S(res))
	require.Equal(t, "abc", res.Hello)
}
