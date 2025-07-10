package data_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/orzkratos/demokratos/demo1kratos"
	"github.com/orzkratos/demokratos/demo1kratos/internal/biz"
	"github.com/orzkratos/demokratos/demo1kratos/internal/data/wiredata"
	"github.com/orzkratos/demokratos/demo1kratos/internal/pkg/cfgdata"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

var caseWireBox *wiredata.WireBox

func TestMain(m *testing.M) {
	configPath := osmustexist.ROOT(filepath.Join(demo1kratos.SourceRoot(), "configs"))
	zaplog.LOG.Debug("run-test-main", zap.String("config-path", configPath))

	cfg := cfgdata.ParseConfig(configPath)
	wireBox, cleanup, err := wiredata.NewWireBox(cfg.Data, log.NewStdLogger(os.Stdout))
	must.Done(err)
	defer cleanup()

	caseWireBox = wireBox
	zaplog.LOG.Debug("run-test-main")
	m.Run()
}

func Test_greeterRepo_Save(t *testing.T) {
	res, err := caseWireBox.GreeterRepo.Save(context.Background(), &biz.Greeter{
		Hello: "xyz",
	})
	require.NoError(t, err)
	t.Log(neatjsons.S(res))
	require.Equal(t, "xyz", res.Hello)
}

func Test_greeterRepo_Update(t *testing.T) {
	res, err := caseWireBox.GreeterRepo.Update(context.Background(), &biz.Greeter{
		Hello: "uvw",
	})
	require.NoError(t, err)
	t.Log(neatjsons.S(res))
	require.Equal(t, "uvw", res.Hello)
}
