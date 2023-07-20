package main

import (
	"api/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", routes.Hello)

	return r
}

// set main branch to track origin/main
// git branch -u origin/main main

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
