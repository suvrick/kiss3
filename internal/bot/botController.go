package bot

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suvrick/kiss/internal/until"
)

type BotController struct {
	service IBotService
}

func NewBotController(service IBotService) *BotController {
	return &BotController{
		service: service,
	}
}

func (ctrl *BotController) Register(router *gin.Engine) {
	botGroup := router.Group("api/bots")
	botGroup.GET("/", ctrl.Get)
	botGroup.PUT("/", ctrl.Update)
	botGroup.POST("/", ctrl.Create)
	botGroup.DELETE("/:id", ctrl.Delete)
}

// curl -X GET localhost:8080/api/bot/
func (ctrl *BotController) Get(c *gin.Context) {

	limit := c.GetInt("limit")

	proxies, err := ctrl.service.Get(c.Request.Context(), limit)

	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Get proxy list fail", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Get proxy list success", nil, proxies)
}

// curl -X POST localhost:8080/api/bot/ --data {\"name\":\"anton\",\"balance\":333,\"game_id\":23423423}
func (ctrl *BotController) Create(c *gin.Context) {

	dto := Bot{}

	err := c.BindJSON(&dto)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Create fail.", err, nil)
		return
	}

	p, err := ctrl.service.Create(c.Request.Context(), dto)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Create fail.", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Create success!", nil, p)
}

// curl -X PUT localhost:8080/api/bot/ --data {\"id\":1,\"balance\":555}
func (ctrl *BotController) Update(c *gin.Context) {

	dto := Bot{}

	err := c.BindJSON(&dto)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Update fail.", err, nil)
		return
	}

	dto, err = ctrl.service.Update(c.Request.Context(), dto)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Update fail.", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Update success!", nil, dto)
}

// curl -X DELETE localhost:8080/api/bot/1
func (ctrl *BotController) Delete(c *gin.Context) {

	ID := c.Param("id")
	botID, err := strconv.Atoi(ID)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Delete fail.", err, nil)
		return
	}

	err = ctrl.service.Delete(c.Request.Context(), uint64(botID))
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Delete fail.", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Delete success!", nil, nil)
}
