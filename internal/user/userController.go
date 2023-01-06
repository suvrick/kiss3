package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/suvrick/kiss/internal/until"
)

type UserController struct {
	service IUserService
}

func NewUserController(service IUserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (ctrl *UserController) Register(router *gin.Engine) {
	userGroup := router.Group("api/users")
	userGroup.GET("/", ctrl.Get)
	userGroup.GET("/user/:id", ctrl.GetByID)
	userGroup.PUT("/", ctrl.Update)
	userGroup.DELETE("/:id", ctrl.Delete)

	auth := router.Group("api/auth")
	auth.POST("/login", ctrl.SingIn)
	auth.POST("/register", ctrl.Create)
}

// curl -X POST localhost:8080/api/users/singin --data {\"email\":\"1@mail.ru\",\"password\":\"123456\"}
func (ctrl *UserController) SingIn(c *gin.Context) {

	type FormData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	dto := FormData{}

	err := c.BindJSON(&dto)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "SingIn fail.", err, nil)
		return
	}

	token, err := ctrl.service.SingIn(c.Request.Context(), dto.Email, dto.Password)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "SingIn fail.", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "SingIn successed!", nil, map[string]interface{}{
		"token": token,
	})
}

// curl -X POST localhost:8080/api/users/ --data {\"email\":\"1@mail.ru\",\"password\":\"123456\",\"password2\":\"123456\"}
func (ctrl *UserController) Create(c *gin.Context) {

	type FormData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	dto := FormData{}

	err := c.BindJSON(&dto)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Create user fail.", err, nil)
		return
	}

	if len(dto.Email) == 0 || len(dto.Password) == 0 {
		until.HTTPResponse(c, http.StatusBadRequest, "Create user fail.", err, nil)
		return
	}

	p, err := ctrl.service.Create(c.Request.Context(), dto.Email, dto.Password)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Create fail.", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Create success!", nil, p)
}

// curl -X GET localhost:8080/api/users/user/1
func (ctrl *UserController) GetByID(c *gin.Context) {

	id := c.Param("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Get user by id fail", err, nil)
		return
	}

	user, err := ctrl.service.GetByID(c.Request.Context(), uint64(userID))

	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Get user by id fail", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Get user by id success", nil, user)
}

// curl -X GET localhost:8080/api/proxy
func (ctrl *UserController) Get(c *gin.Context) {

	limit := c.GetInt("limit")

	users, err := ctrl.service.Get(c.Request.Context(), limit)

	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Get user list fail", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Get user list success", nil, users)
}

// curl -X PUT localhost:8080/api/proxy/update --data {\"id\":2,\"is_bad\":true}
func (ctrl *UserController) Update(c *gin.Context) {

	dto := User{}

	err := c.BindJSON(&dto)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Update fail.", err, nil)
		return
	}

	dto, err = ctrl.service.Update(c.Request.Context(), dto)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Create fail.", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Create success!", nil, dto)
}

// curl -X DELETE localhost:8080/api/proxy/delete/1
func (ctrl *UserController) Delete(c *gin.Context) {

	ID := c.Param("id")
	proxyID, err := strconv.Atoi(ID)
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Delete fail.", err, nil)
		return
	}

	err = ctrl.service.Delete(c.Request.Context(), uint64(proxyID))
	if err != nil {
		until.HTTPResponse(c, http.StatusBadRequest, "Delete fail.", err, nil)
		return
	}

	until.HTTPResponse(c, http.StatusOK, "Delete success!", nil, nil)
}
