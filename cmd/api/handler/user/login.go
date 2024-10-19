package handler

import (
	"goWallet/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) Login(c *gin.Context) {
	var userLogin *domain.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	token, err := h.User.Login(userLogin)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"token": token})

}
