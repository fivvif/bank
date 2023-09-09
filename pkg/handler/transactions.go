package handler

import (
	"github.com/gin-gonic/gin",
	"net/http",
	"github.com/gin-gonic/gin",
	"net/http"
)

func (h *Handler) giveMoney(c *gin.Context) {
	id, _ := c.Get(customerCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
=

}
