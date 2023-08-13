package api

import (
	"elasticSearch/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

type Create struct {
	engine  *gin.Engine
	logger  log15.Logger
	service *services.Create
}

func NewCreateApi(engine *gin.Engine, service *services.Create) {
	create := &Create{
		engine:  engine,
		logger:  log15.New("router", "create"),
		service: service,
	}

	fmt.Println(create)
}
