package e2etest

import (
	"testing"

	"github.com/lightningnetwork/lnd/signal"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/dapplink-labs/bbn-fp-l2/eotsmanager"
	"github.com/dapplink-labs/bbn-fp-l2/eotsmanager/config"
	"github.com/dapplink-labs/bbn-fp-l2/eotsmanager/service"
)

type EOTSServerHandler struct {
	t           *testing.T
	interceptor *signal.Interceptor
	eotsServer  *service.Server
	Cfg         *config.Config
}

func NewEOTSServerHandler(t *testing.T, cfg *config.Config, eotsHomeDir string, shutdownInterceptor signal.Interceptor) *EOTSServerHandler {
	dbBackend, err := cfg.DatabaseConfig.GetDBBackend()
	require.NoError(t, err)
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	logger, err := loggerConfig.Build()
	require.NoError(t, err)
	eotsManager, err := eotsmanager.NewLocalEOTSManager(eotsHomeDir, cfg.KeyringBackend, dbBackend, logger)
	require.NoError(t, err)

	eotsServer := service.NewEOTSManagerServer(cfg, logger, eotsManager, dbBackend, shutdownInterceptor)

	return &EOTSServerHandler{
		t:           t,
		eotsServer:  eotsServer,
		interceptor: &shutdownInterceptor,
		Cfg:         cfg,
	}
}

func (eh *EOTSServerHandler) Start() {
	go eh.startServer()
}

func (eh *EOTSServerHandler) startServer() {
	err := eh.eotsServer.RunUntilShutdown()
	require.NoError(eh.t, err)
}

func (eh *EOTSServerHandler) Stop() {
	eh.interceptor.RequestShutdown()
}