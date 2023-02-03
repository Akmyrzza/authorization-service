package rest

import (
	"net/http"

	"github.com/Akmyrzza/authorization-service/internal/domain/usecases"
	"github.com/gin-gonic/gin"
)

type St struct {
	ucs *usecases.St
}

func GetHandler(ucs *usecases.St) http.Handler {
	r := gin.New()

	s := &St{ucs: ucs}
	// user
	r.POST("/user", s.hUserCreate)

	return r
}

//func (o *St) getRequestContext(c *gin.Context) context.Context {
//	return o.ucs.
//}
