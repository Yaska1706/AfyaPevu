package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaska1706/AfyaPevu/internal/pkg/persistence"
	"github.com/yaska1706/AfyaPevu/pkg/crypto"
	"github.com/yaska1706/AfyaPevu/pkg/http-err"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var loginInput LoginInput
	_ = c.BindJSON(&loginInput)
	s := persistence.GetUserRepository()
	if user, err := s.GetByUsername(loginInput.Username); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if !crypto.ComparePasswords(user.Hash, []byte(loginInput.Password)) {
			http_err.NewError(c, http.StatusForbidden, errors.New("user and password not match"))
			return
		}
		token, _ := crypto.CreateToken(user.Username)
		c.JSON(http.StatusOK, token)
	}
}
