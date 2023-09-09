package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	customerCtx         = "customerId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty Authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid Authorization header")
		return
	}

	customerId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(customerCtx, customerId)

}
