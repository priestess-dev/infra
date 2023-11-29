package oauth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server interface {
	GinLoginHandler(c *gin.Context)
	GinCallbackHandler(c *gin.Context)
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleCallback(w http.ResponseWriter, r *http.Request)
}
