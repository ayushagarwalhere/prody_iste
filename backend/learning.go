package main // how to run a server in go?

import (
  "github.com/gin-gonic/gin" // import the gin framework
  "net/http" // import the net/http package for handling HTTP requests
)

func main() {
  router := gin.Default() // creates a default router with logger and recovery middleware
  router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{ //http.StatusOK is a constant for 200 status code, gin.H is a shortcut for map[string]interface{}
      "message": "pong",
    })
  })
  router.Run() // default port is 8080 can also set by passing ":8080" as an argument
}