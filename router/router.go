package router

import (
	"elasticSearch/config"
	"elasticSearch/router/api"
	"elasticSearch/services"
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
	service *services.ServiceRoot
}

func NewRouter(cfg *config.Config, service *services.ServiceRoot) (*Router, error) {
	r := &Router{
		engine:  gin.New(),
		config:  cfg,
		logger:  log15.New("module", "api"),
		service: service,
	}

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

	api.NewCreateApi(r.engine, r.service.Create)
	api.NewSearchApi(r.engine, r.service.Search)
	api.NewUpdateApi(r.engine, r.service.Update)
	api.NewDeleteApi(r.engine, r.service.Delete)

	return r, nil
}

func (r *Router) Run(port string) error {
	msg := fmt.Sprintf("Server Started In %v", port)
	r.logger.Info(msg, "info", nil)
	return r.engine.Run(port)
}
