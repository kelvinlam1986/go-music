package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-music/repositories"
	"go-music/utils"
	"net/http"
	"time"
)

const SECRET_KEY = "kelvinlam_1986"

type AccountController struct {
	AccountRepository repositories.IAccountRepository
}

func NewAccountController(repo repositories.IAccountRepository) *AccountController {
	return &AccountController{ AccountRepository: repo }
}

func (controller *AccountController) SignIn(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	account, err := controller.AccountRepository.GetAccountByUserName(username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{ "error": "Username and password not match" })
		return
	}

	valid :=  utils.CheckPassword(account.Password, password)
	if !valid {
		context.JSON(http.StatusForbidden, gin.H{"error": "Username and password not match"})
		return
	}

	claims := jwt.MapClaims{
		"username": username,
		"ExpiresAt": 15000,
		"IssuedAt": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	bytedSecret := []byte(SECRET_KEY)
	tokenString, err := token.SignedString(bytedSecret)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{ "error": err.Error() })
		return
	}

	context.JSON(http.StatusOK, gin.H{ "token": tokenString, "status": "success" })
}