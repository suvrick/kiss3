package user

import (
	"log"
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
	userGroup.GET("/:id", ctrl.GetByID)
	userGroup.PUT("/", ctrl.Update)
	userGroup.DELETE("/:id", ctrl.Delete)

	auth := router.Group("api/auth")
	auth.POST("/login", ctrl.Login)
	auth.POST("/register", ctrl.Create)
}

// curl -X POST localhost:8080/api/users/singin --data {\"email\":\"1@mail.ru\",\"password\":\"123456\"}
func (ctrl *UserController) Login(c *gin.Context) {

	type FormData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	dto := FormData{}

	err := c.BindJSON(&dto)
	if err != nil {
		log.Printf("[Login] Login fail. %v\n", err)
		until.HTTPResponse(c, http.StatusBadRequest, "Login fail.", nil, nil)
		return
	}

	token, err := ctrl.service.SingIn(c.Request.Context(), dto.Email, dto.Password)
	if err != nil {
		log.Printf("[Login] Create token fail. %v\n", err)
		until.HTTPResponse(c, http.StatusBadRequest, "Login fail.", nil, nil)
		return
	}

	log.Printf("[Login] Login user successed. %v\n", dto)
	until.HTTPResponse(c, http.StatusOK, "Login successed!", nil, map[string]interface{}{
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
		log.Printf("[Create] Create user fail. %v\n", err.Error())
		until.HTTPResponse(c, http.StatusBadRequest, "Create user fail.", nil, nil)
		return
	}

	if len(dto.Email) == 0 || len(dto.Password) == 0 {
		log.Println("[Create] Create user fail. Empty password or email")
		until.HTTPResponse(c, http.StatusBadRequest, "Create user fail.", nil, nil)
		return
	}

	log.Printf("[Create] Try create new user: %v\n", dto)
	user, err := ctrl.service.Create(c.Request.Context(), dto.Email, dto.Password)
	if err != nil {
		log.Printf("[Create] Create user fail. %s\n", err.Error())
		until.HTTPResponse(c, http.StatusBadRequest, "Create user fail.", nil, nil)
		return
	}

	log.Printf("[Create] Create user successed. %v\n", user)
	until.HTTPResponse(c, http.StatusOK, "Create user success!", nil, user)
}

// curl -X GET localhost:8080/api/users/user/1
func (ctrl *UserController) GetByID(c *gin.Context) {

	id := c.Param("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[GetByID] Get user by id fail. %v\n", err.Error())
		until.HTTPResponse(c, http.StatusBadRequest, "Get user by id fail", nil, nil)
		return
	}

	user, err := ctrl.service.GetByID(c.Request.Context(), uint64(userID))

	if err != nil {
		log.Printf("[GetByID] Get user by id fail. %v\n", err.Error())
		until.HTTPResponse(c, http.StatusBadRequest, "Get user by id fail", nil, nil)
		return
	}

	log.Printf("[GetByID] Get user by id successed. %v\n", user)
	until.HTTPResponse(c, http.StatusOK, "Get user by id success", nil, user)
}

// curl -X GET localhost:8080/api/proxy
func (ctrl *UserController) Get(c *gin.Context) {

	limit := c.GetInt("limit")

	users, err := ctrl.service.Get(c.Request.Context(), limit)

	if err != nil {
		log.Printf("[Get] Get users list fail. %v\n", err.Error())
		until.HTTPResponse(c, http.StatusBadRequest, "Get user list fail", nil, nil)
		return
	}

	log.Printf("[Get] Get users list success. Users count %d\n", len(users))
	until.HTTPResponse(c, http.StatusOK, "Get user list success", nil, users)
}

// curl -X PUT localhost:8080/api/proxy/update --data {\"id\":2,\"is_bad\":true}
func (ctrl *UserController) Update(c *gin.Context) {

	dto := User{}

	err := c.BindJSON(&dto)
	if err != nil {
		log.Printf("[Update] Update user fail. %v, %v\n", err.Error(), dto)
		until.HTTPResponse(c, http.StatusBadRequest, "Update user fail.", nil, nil)
		return
	}

	dto, err = ctrl.service.Update(c.Request.Context(), dto)
	if err != nil {
		log.Printf("[Update] Update user fail. %v, %v\n", err.Error(), dto)
		until.HTTPResponse(c, http.StatusBadRequest, "Update user fail.", nil, nil)
		return
	}

	log.Printf("[Update] Update user successed. %v\n", dto)
	until.HTTPResponse(c, http.StatusOK, "Update user successed!", nil, dto)
}

// curl -X DELETE localhost:8080/api/proxy/delete/1
func (ctrl *UserController) Delete(c *gin.Context) {

	ID := c.Param("id")
	userID, err := strconv.Atoi(ID)
	if err != nil {
		log.Printf("[Delete] Delete user fail. %v, %v\n", err.Error(), userID)
		until.HTTPResponse(c, http.StatusBadRequest, "Delete user fail.", nil, nil)
		return
	}

	err = ctrl.service.Delete(c.Request.Context(), uint64(userID))
	if err != nil {
		log.Printf("[Delete] Delete user fail. %v, %v\n", err.Error(), userID)
		until.HTTPResponse(c, http.StatusBadRequest, "Delete user fail.", nil, nil)
		return
	}

	log.Printf("[Delete] Delete user successed. %v\n", userID)
	until.HTTPResponse(c, http.StatusOK, "Delete successed!", nil, nil)
}
