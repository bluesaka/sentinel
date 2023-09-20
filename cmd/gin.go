package main

import (
	"github.com/alibaba/sentinel-golang/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sentinel"
)

func main() {
	sentinel.InitSentinel()
	r := gin.Default()
	r.GET("/order/get", handleGetOrder)
	r.Run(":8080")
	// 模拟请求
	// ab -n 100 -c 20 http://localhost:8080/order/get
}

func handleGetOrder(c *gin.Context) {
	e, err := api.Entry("/order/get")
	if err != nil {
		// Blocked. We could get the block reason from the BlockError.
		c.JSON(http.StatusTooManyRequests, gin.H{"msg": err.BlockMsg()})
		return
	} else {
		// Passed, wrap the logic here.
		//time.Sleep(time.Millisecond)
		log.Println("pass")
		// Be sure the entry is exited finally.
		e.Exit()
	}

}
