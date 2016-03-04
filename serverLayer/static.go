package main

import (
  "github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
    c.Next()
  }
}
  
func main() {
  router := gin.Default()
  router.Use(Cors())

  router.Static("/app", "../clientLayer")
  //router.StaticFS("/more_static", http.Dir("my_file_system"))
  //router.StaticFile("/favicon.ico", "./resources/favicon.ico")

  // Listen and server on 0.0.0.0:8080
  router.Run(":8000")

}