package rest

import (
	"github.com/Akmyrzza/authorization-service/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

func (o *St) hUserCreate(c *gin.Context) {
	reqObj := &entities.User{}
	if err := c.BindJSON(reqObj); err != nil {
		return
	}

	result, err := o.ucs.UserCreate()
}