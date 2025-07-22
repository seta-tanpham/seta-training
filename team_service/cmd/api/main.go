package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	Connect()
	r := gin.Default()

	api := r.Group("")
	{
		api.POST("/teams", CreateTeam)
		api.POST("/teams/:teamId/members", AddMember)
		api.DELETE("/teams/:teamId/members/:memberId", RemoveMember)
		api.POST("/teams/:teamId/managers", AddManager)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Teams API is running",
		})
	})

	r.Run(":8080")
}
