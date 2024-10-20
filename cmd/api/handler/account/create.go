package account

import (
	"goWallet/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) Create(c *gin.Context) {
	var account *domain.Account

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Account.Create(account); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": &account})
}
