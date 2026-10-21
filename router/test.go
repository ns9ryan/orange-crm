package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"orange-crm/utils"
	"os"
)

type Test struct{}

func (g *Test) Group(r *gin.RouterGroup) {
	{
		// 健康监测
		r.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, utils.TT("ok"))
		})
	}

	{
		r.GET("/hello/:name", func(c *gin.Context) {
			var msg = c.Param("name")
			c.JSON(http.StatusOK, gin.H{
				"msg": "hello " + msg,
			})
		})
	}

	{
		r.GET("/test", func(c *gin.Context) {
			// 模拟一个 network operation 错误
			err := &net.OpError{
				Op:     "dial",
				Net:    "tcp",
				Source: nil,
				Addr:   nil,
				//Err:    fmt.Errorf("connection refused"),
				// 将 Err 字段设置为一个 *os.SyscallError 错误
				Err: &os.SyscallError{
					Err:     fmt.Errorf("connection reset by peer"),
					Syscall: "broken pipe",
				},
			}
			// 通过 panic 抛出 net.OpError 错误
			panic(err)

			//tmp := []int{1, 2}
			//fmt.Println(tmp[3])
			//panic("致命错误")
		})
	}
}
