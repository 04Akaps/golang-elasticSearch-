package apps

import (
	"elasticSearch/config"
	"elasticSearch/router"
	"github.com/inconshreveable/log15"
	"os"
)

type App struct {
	config *config.Config
	router *router.Router
	logger log15.Logger
	stop   chan struct{}
}

func NewApps(cfg *config.Config) *App {
	app := &App{
		config: cfg,
		logger: log15.New("module", "app"),
		stop:   make(chan struct{}),
	}
	var err error

	if app.router, err = router.NewRouter(cfg); err != nil {
		app.logger.Crit("Server NewRouter Error", "crit", err)
		os.Exit(0)
	}

	if err = app.router.Run(cfg.Server.Port); err != nil {
		app.logger.Crit("Server Start Error", "Crit", err)
		os.Exit(0)
	}

	return app
}

func (a *App) Wait() {
	a.logger.Info("Starting Server..")
	<-a.stop
}
