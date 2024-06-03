package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"imdb_project/data/dto"
	"imdb_project/service"
	"net/http"
)

type UserController struct {
	UserService service.IUserService
	Validate    *validator.Validate
	Store       *sessions.CookieStore
}

func (c *UserController) SubscribeEndpoints(engine *gin.RouterGroup) {
	engine.POST("/user", c.CreateUser)
	engine.GET("/user/id", c.FindUserById)
	engine.GET("/user/all", c.FindAllUsers)
	engine.GET("/user/email", c.FindUserByEmail)
}

func NewUserController(userService service.IUserService, validator *validator.Validate, store *sessions.CookieStore) *UserController {
	return &UserController{UserService: userService, Validate: validator, Store: store}
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
	userID := c.getUserID(ctx)
	response := c.UserService.FindUserById(userID.String())
	ctx.JSON(int(response.StatusCode), response)
}

func (c *UserController) FindAllUsers(ctx *gin.Context) {
	response := c.UserService.FindAllUsers()
	ctx.JSON(int(response.StatusCode), response)
}

func (c *UserController) FindUserByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	response := c.UserService.FindUserByEmail(email)
	ctx.JSON(int(response.StatusCode), response)
}

// ...
func (c *UserController) getUserID(ctx *gin.Context) uuid.UUID {
	session, err := c.Store.Get(ctx.Request, "imdb-session")
	if err != nil {
		return uuid.UUID{}
	}

	userID, _ := session.Values["id"]

	return uuid.MustParse(userID.(string))
}
