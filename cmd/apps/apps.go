package apps

import (
	"context"
	"elasticSearch/config"
	"elasticSearch/router"
	"elasticSearch/services"
	"elasticSearch/types"
	"encoding/json"
	"github.com/inconshreveable/log15"
	"github.com/olivere/elastic/v7"
	"io"
	"os"
	"strconv"
	"time"
)

type App struct {
	config  *config.Config
	router  *router.Router
	logger  log15.Logger
	service *services.ServiceRoot

	stop chan struct{}
}

func NewApps(cfg *config.Config, migration bool) *App {
	app := &App{
		config: cfg,
		logger: log15.New("module", "app"),
		stop:   make(chan struct{}),
	}

	var err error

	if app.service, err = services.NewService(cfg); err != nil {
		app.logger.Crit("Server NewService Error", "crit", err)
		os.Exit(0)
	}

	if migration {
		if file, err := os.Open("./migration.json"); err != nil {
			os.Exit(0)
		} else {
			defer file.Close()

			if fileContent, err := io.ReadAll(file); err != nil {
				app.logger.Crit("can't read file")
				os.Exit(0)
			} else {
				var m []types.MigrationFile

				if err = json.Unmarshal(fileContent, &m); err != nil {
					app.logger.Crit("failed to unMarshal File", "err", err)
					os.Exit(0)
				} else {
					app.logger.Info("migration started", "time", time.Now().Unix())
					elasticClient := app.service.ElasticSearch.Client

					bulk := elasticClient.Bulk()

					for _, data := range m {
						var m types.Migration

						m.KeyWord = data.KeyWord
						m.Owner = data.Owner
						m.PK = data.ObjectId.Hex()

						if value, ok := data.Tid.(map[string]interface{}); ok {
							// numberLong타입인 경우
							tokenId := value["$numberLong"]
							m.Tid, _ = strconv.Atoi(tokenId.(string))
						} else if value, ok := data.Tid.(float64); ok {
							m.Tid = int(int64(value))
						} else {
							m.Tid = data.Tid.(int)
						}

						model := elastic.NewBulkUpdateRequest()
						model.Id(m.PK)
						model.Index(types.Index)
						model.Doc(&m)
						model.DocAsUpsert(true)

						bulk.Add(model)
					}

					if response, err := bulk.Do(context.Background()); err != nil {
						app.logger.Crit("failed to bulk Migration")
						os.Exit(0)
					} else {
						successedItem := len(response.Succeeded())
						updatedItem := len(response.Updated())
						app.logger.Info("success to migration", "success", successedItem, "updated", updatedItem)
						os.Exit(0)
					}

				}

			}

		}
	}

	if app.router, err = router.NewRouter(cfg, app.service); err != nil {
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

func (a *App) getMigrationJson() {

}
