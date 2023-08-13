package router

import (
	"elasticSearch/config"
	"elasticSearch/repository"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
	"time"
)

type Router struct {
	engine  *gin.Engine
	logger  log15.Logger
	config  *config.Config
	elastic *repository.Elastic
}

func NewRouter(cfg *config.Config) (*Router, error) {
	r := &Router{
		engine: gin.New(),
		config: cfg,
		logger: log15.New("module", "router"),
	}
	var err error

	r.engine.Use(gin.Logger())
	r.engine.Use(gin.Recovery())
	r.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		ExposeHeaders:    []string{"ORIGIN", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	if r.elastic, err = repository.NewElastic(cfg); err != nil {
		r.logger.Crit("Failed Connect ElasticSearch", "crit", err)
		return nil, err
	}

	return r, nil
}

func (r *Router) Run(port string) error {
	msg := fmt.Sprintf("Server Started In %v", port)
	r.logger.Info(msg, "info", nil)
	return r.engine.Run(port)
}
