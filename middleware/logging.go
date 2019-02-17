package middleware

import (
	"github.com/gin-gonic/gin"
	"bytes"
	"time"
	"regexp"
	"io/ioutil"
	"encoding/json"
	"github.com/lexkong/log"
	"github.com/willf/pad"
	"myapiserver/pkg/result"
	"myapiserver/pkg/errno"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error)  {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc  {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		reg := regexp.MustCompile("/v1|/login")
		if !reg.MatchString(path) {
			return
		}

		var bodyBytes []byte

		// Read the Body content
		if c.Request.Body != nil {

			/**
				ReadAll 读取 r 中的所有数据，返回读取的数据和遇到的错误。
				如果读取成功，则 err 返回 nil，而不是 EOF，因为 ReadAll 定义为读取
				所有数据，所以不会把 EOF 当做错误处理。
			 */
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// NopCloser 将 r 包装为一个 ReadCloser 类型，但 Close 方法不做任何事情。
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		method := c.Request.Method
		ip := c.ClientIP()

		blw := &bodyLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}

		c.Writer = blw

		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""

		// get code and message
		var response result.Response
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			log.Errorf(err, "response body can not unmarshal to model.Response struct, body: `%s`", blw.body.Bytes())
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}

		log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message)

	}
}