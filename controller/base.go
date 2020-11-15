package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回数据格式拼装
type Context struct {
	*gin.Context
}

const CodeOk = 0
const CodeParamError = -1
const CodeServerError = 500

func (c *Context) success(data interface{}) {
	if data == nil {
		data = map[string]string{}
	}
	c.JSON(http.StatusOK, gin.H{
		"rid":  c.MustGet("rid").(string),
		"code": CodeOk,
		"msg":  "ok",
		"data": data,
	})
}
func (c *Context) error(code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"rid":  c.MustGet("rid").(string),
		"code": code,
		"msg":  msg,
	})
}