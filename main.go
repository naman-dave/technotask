package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/users", HandleUsers)
	r.GET("/user/:id", HandleUserByID)
	r.POST("/user", HandleAddUser)
	r.PUT("/user/:id", HandleEditUsers)
	r.DELETE("/user/:id", HandleDeleteUsers)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
