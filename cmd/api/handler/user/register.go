package handlerUser

import (
	"goWallet/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) Register(c *gin.Context) {
	var user domain.User
	var account domain.Account
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.User.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	account.UserID = user.ID

	if err := h.Account.Create(&account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
