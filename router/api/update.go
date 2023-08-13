package api

import (
	"elasticSearch/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

type Update struct {
	engine  *gin.Engine
	logger  log15.Logger
	service *services.Update
}

func NewUpdateApi(engine *gin.Engine, service *services.Update) {
	update := &Update{
		engine:  engine,
		logger:  log15.New("router", "update"),
		service: service,
	}

	fmt.Println(update)

}
