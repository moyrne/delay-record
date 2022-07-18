package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moyrne/delay-record/pkg/pingserver/internal/biz"
)

func Run() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		resp, err := biz.Ping(c.Request.Context())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, c.Error(err))
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
