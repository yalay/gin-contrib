package middlewares

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var methodsForModify = map[string]bool{
	http.MethodPost:   true,
	http.MethodPut:    true,
	http.MethodPatch:  true,
	http.MethodDelete: true,
}

const (
	ContextKeyReqBody = "req_body"
)

// store request body to context, so other handlers can use it directly
// only for Post,Put,Patch,Delete method
func StoreRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !methodsForModify[c.Request.Method] {
			return
		}

		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		rewriteBody := ioutil.NopCloser(bytes.NewBuffer(reqBody))

		c.Request.Body = rewriteBody
		c.Set(ContextKeyReqBody, string(reqBody))

		c.Next()
	}
}
