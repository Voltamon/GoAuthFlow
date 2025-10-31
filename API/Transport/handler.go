package transport

import (
	"net/http"

	data "github.com/Voltamon/GoAuthFlow/API/Data"
	auth "github.com/Voltamon/GoAuthFlow/API/Service"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Handler struct {
	authService *auth.AuthService
}

func NewHandler(authService *auth.AuthService) *Handler {
	return &Handler{authService: authService}
}

func (h *Handler) Register(c *gin.Context) {
	var info UserInfo

	err := c.ShouldBindJSON(&info)
	if err != nil {
		auth.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	user := data.User{
		Hash: auth.GenerateHash(info.Email, info.Password),
	}
	u, err := h.authService.RegisterUser(&user)
	if err != nil {
		auth.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c *gin.Context) {
	var info UserInfo

	err := c.ShouldBindJSON(&info)
	if err != nil {
		auth.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	user := data.User{
		Hash: auth.GenerateHash(info.Email, info.Password),
	}

	u, err := h.authService.LoginUser(&user)
	if err != nil {
		auth.ResponseError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, u)
}
