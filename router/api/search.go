package api

import (
	"elasticSearch/services"
	"elasticSearch/types"
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

	baseUri := "search"

	search.engine.POST(baseUri+"/user-by-name/:name", search.searchUserByName)
	search.engine.POST(baseUri+"/user-by-age/:age", search.searchUserByAge)
	search.engine.POST(baseUri+"/user-by-address/:address", search.searchUserByAddress)

}

func (m *Search) searchUserByName(c *gin.Context) {
	var req types.SearchByNameReq

	if err := c.ShouldBindUri(&req); err != nil {
		errResponse(c, err.Error())
	} else if err = c.ShouldBindJSON(&req); err != nil {
		errResponse(c, err.Error())
	} else if response, err := m.service.SearchByName("user", req.Name, req.Size, req.Sort); err != nil {
		errResponse(c, err.Error())
	} else {
		successResponse(c, response)
	}
}

func (m *Search) searchUserByAge(c *gin.Context) {
	var req types.SearchByAgeReq

	if err := c.ShouldBindUri(&req); err != nil {
		errResponse(c, err.Error())
	} else if err = c.ShouldBindJSON(&req); err != nil {
		errResponse(c, err.Error())
	} else if response, err := m.service.SearchByAge("user", req.Age, req.Size, req.Sort); err != nil {
		errResponse(c, err.Error())
	} else {
		successResponse(c, response)
	}
}

func (m *Search) searchUserByAddress(c *gin.Context) {
	var req types.SearchByAddressReq

	if err := c.ShouldBindUri(&req); err != nil {
		errResponse(c, err.Error())
	} else if err = c.ShouldBindJSON(&req); err != nil {
		errResponse(c, err.Error())
	} else if response, err := m.service.SearchByAddress("user", req.Address, req.Size, req.Sort); err != nil {
		errResponse(c, err.Error())
	} else {
		successResponse(c, response)
	}
}
