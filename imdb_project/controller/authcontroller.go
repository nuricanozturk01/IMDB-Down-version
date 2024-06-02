package controller

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"imdb_project/data/dto"
	"imdb_project/service"
	"imdb_project/util"
	"io"
	"log"
	"net/http"
	"os"
)

type AuthController struct {
	AuthenticationService *service.AuthenticationService
	InformationService    *service.InformationService
	Validate              *validator.Validate
	Store                 *sessions.CookieStore
}

var googleOauthConfig *oauth2.Config

func NewAuthController(authService *service.AuthenticationService, informationService *service.InformationService, validator *validator.Validate, store *sessions.CookieStore) *AuthController {
	return &AuthController{AuthenticationService: authService, InformationService: informationService, Validate: validator, Store: store}
}

func (controller *AuthController) SubscribeEndpoints(engine *gin.Engine) {
	engine.POST("/api/auth/login", controller.Login)
	engine.POST("/api/auth/register", controller.Register)
	engine.GET("/api/auth/google/login", controller.GoogleLogin)
	engine.GET("/api/auth/google/callback", controller.GoogleCallback)
	engine.GET("/api/countries/all", controller.FindAllCountries)
	engine.GET("/api/city/by-country", controller.FindCitiesByCountry)
}

func init() {
	clientID := os.Getenv("CLIENT_ID")
	clientSECRETS := os.Getenv("CLIENT_KEY")
	redirectURL := os.Getenv("REDIRECT_URL")

	if clientID == "" || clientSECRETS == "" || redirectURL == "" {
		log.Fatal("Missing required environment variables CLIENT_ID, CLIENT_SECRET, or REDIRECT_URL")
	}

	googleOauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSECRETS,
		RedirectURL:  redirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func (controller *AuthController) GoogleLogin(ctx *gin.Context) {
	url := googleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (controller *AuthController) GoogleCallback(ctx *gin.Context) {
	code := ctx.Query("code")

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token: " + err.Error()})
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info: " + err.Error()})
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	userInfo, err := io.ReadAll(response.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read user info: " + err.Error()})
		return
	}

	var googleUserDTO dto.GoogleUserDTO

	// Deserialize the user info
	if err := json.Unmarshal(userInfo, &googleUserDTO); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal user info: " + err.Error()})
		return
	}

	result := controller.AuthenticationService.LoginOAuth2(&googleUserDTO)

	if result.StatusCode == 201 || result.StatusCode == 200 {
		session, err := controller.Store.Get(ctx.Request, "imdb-session")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get session: " + err.Error()})
			return
		}
		session.Values["authenticated"] = true
		session.Values["user"] = googleUserDTO.Email
		session.Values["id"] = result.Data.ID.String()

		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
			Secure:   false,
		}

		err = session.Save(ctx.Request, ctx.Writer)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session: " + err.Error()})
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:4200")
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log in with OAuth2"})
	}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	var login *dto.LoginDTO
	err := ctx.BindJSON(&login)

	if validationErr := controller.Validate.Struct(login); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := controller.AuthenticationService.Login(login.Email, login.Password)
	if result.StatusCode == 201 || result.StatusCode == 200 {
		session, _ := controller.Store.Get(ctx.Request, "imdb-session")
		session.Values["authenticated"] = true
		session.Values["user"] = login.Email
		session.Values["id"] = result.Data.ID.String()

		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
			Secure:   false,
		}

		err := session.Save(ctx.Request, ctx.Writer)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session: " + err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Logged in successfully", "sessionID": session.ID})
	}
	ctx.JSON(int(result.StatusCode), result)
}

func (controller *AuthController) Register(ctx *gin.Context) {
	var register *dto.UserCreateDTO

	err := ctx.BindJSON(&register)

	if validationErr := controller.Validate.Struct(register); validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pass, _ := util.HashPassword(register.Password)
	register.Password = pass
	response := controller.AuthenticationService.Register(register)

	ctx.JSON(int(response.StatusCode), response)
}

func (controller *AuthController) Logout(ctx *gin.Context) {
	session, _ := controller.Store.Get(ctx.Request, "imdb-session")
	session.Options.MaxAge = -1
	err := session.Save(ctx.Request, ctx.Writer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (controller *AuthController) FindAllCountries(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"countries": controller.InformationService.FindAllCountries()})
}

func (controller *AuthController) FindCitiesByCountry(c *gin.Context) {
	country := c.Query("country")
	c.JSON(http.StatusOK, gin.H{"cities": controller.InformationService.FindCitiesByCountry(country)})
}
