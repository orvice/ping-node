package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrResponse struct {
	Code    int
	Message string
}

func homeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"cat": "memo",
	})
}

type PingRequest struct {
	Addr string `json:"addr"`
}

func pingHandler(c *gin.Context) {
	var req PingRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrResponse{
			Code: RequestInvalid,
		})
		return
	}

	static, err := Ping(req.Addr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrResponse{
			Code: PingFailed,
		})
		return
	}
	c.JSON(http.StatusOK, static)
}

func web() {
	r := gin.Default()
	r.GET("/", homeHandler)
	r.POST("/ping", pingHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
