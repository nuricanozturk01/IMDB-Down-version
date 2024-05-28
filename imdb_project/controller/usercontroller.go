package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"imdb_project/data/dto"
	"imdb_project/service"
	"net/http"
)

type UserController struct {
	UserService service.IUserService
	Validate    *validator.Validate
}

func (c *UserController) SubscribeEndpoints(engine *gin.RouterGroup) {
	engine.POST("/api/v1/user", c.CreateUser)
	engine.GET("/api/v1/user/id", c.FindUserById)
	engine.GET("/api/v1/user/all", c.FindAllUsers)
	engine.GET("/api/v1/user/username", c.FindUserByUsername)
	engine.GET("/api/v1/user/email", c.FindUserByEmail)
}

func NewUserController(userService service.IUserService, validator *validator.Validate) *UserController {
	return &UserController{UserService: userService, Validate: validator}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user *dto.UserCreateDTO

	err := ctx.BindJSON(&user)

	if validationErr := c.Validate.Struct(user); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := c.UserService.CreateUser(user)

	ctx.JSON(int(response.StatusCode), response)
}

func (c *UserController) FindUserById(ctx *gin.Context) {
	userID := ctx.Query("id")
	response := c.UserService.FindUserById(userID)
	ctx.JSON(int(response.StatusCode), response)
}

func (c *UserController) FindAllUsers(ctx *gin.Context) {
	response := c.UserService.FindAllUsers()
	ctx.JSON(int(response.StatusCode), response)
}

func (c *UserController) FindUserByUsername(ctx *gin.Context) {
	username := ctx.Query("username")
	response := c.UserService.FindUserByUsername(username)
	ctx.JSON(int(response.StatusCode), response)
}

func (c *UserController) FindUserByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	response := c.UserService.FindUserByEmail(email)
	ctx.JSON(int(response.StatusCode), response)
}

// ...
