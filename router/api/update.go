package api

import (
	"elasticSearch/services"
	"elasticSearch/types"
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

	baseUri := "update"

	update.engine.POST(baseUri+"/update-example/:name/:age", update.updateExample)

}

func (u *Update) updateExample(c *gin.Context) {

	var req types.UpdateUserReq

	if err := c.ShouldBindUri(&req); err != nil {
		errResponse(c, err.Error())
	} else if err = u.service.UpDateUser(types.Index, req.Name, req.Age); err != nil {
		errResponse(c, err.Error())
	} else {
		msg := fmt.Sprintf("success update User Name : %s, age : %d", req.Name, req.Age)
		successResponse(c, msg)
	}
}
