package api

import (
	"elasticSearch/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

type Search struct {
	engine  *gin.Engine
	logger  log15.Logger
	service *services.Search
}

func NewSearchApi(engine *gin.Engine, service *services.Search) {
	search := &Search{
		engine:  engine,
		logger:  log15.New("router", "search"),
		service: service,
	}

	fmt.Println(search)
}
