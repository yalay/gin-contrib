package middlewares

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
)

type LogResponseWriter struct {
	gin.ResponseWriter
	rspBody *bytes.Buffer
}

type logger interface {
	Infof(format string, args ...interface{})
}

func (r *LogResponseWriter) Write(p []byte) (int, error) {
	r.rspBody.Write(p)
	return r.ResponseWriter.Write(p)
}

// access logger with request body and response body
// should be called after StoreRequest middleware
func AccessLogger(logger logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var startTime = time.Now()
		logRspWriter := &LogResponseWriter{c.Writer, bytes.NewBuffer([]byte{})}
		c.Writer = logRspWriter

		c.Next()

		reqBody := c.GetString(ContextKeyReqBody)
		elapsed := float64(time.Now().Sub(startTime).Nanoseconds()) / 1e6
		logger.Infof("%s %s %s %.3fms %s %s %s", c.Request.Method, c.Request.URL.String(),
			logRspWriter.Status(), elapsed, c.ClientIP(), reqBody, logRspWriter.rspBody)
	}
}
