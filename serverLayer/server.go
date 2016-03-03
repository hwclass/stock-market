package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

type Stock struct {
  Id int64 `db:"id" json:"id"`
  Firm string `db:"firm" json:"firm"`
  Data string `db:"data" json:"data"`
} 

func GetStocks(c *gin.Context) {
  c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
  c.Writer.Header().Add("content-type", "*")
  c.Writer.Header().Add("cache-control", "no-cache")
  c.Writer.Header().Add("connection", "keep-alive")
  type Stocks []Stock
  var stocks = Stocks{
    Stock{Firm: "facebook", Data: "Oliver"},
    Stock{Firm: "apple", Data: "Malcom"},
  }
  c.Next()
  c.JSON(200, stocks)
}

func Cors() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
    c.Next()
  }
}

func main() {

  /*
  http.Handle("/", http.FileServer(http.Dir("../clientLayer")))
  http.ListenAndServe(":3000", nil)
  */

  r.Use(Cors())

  v1 := r.Group("api/v1")
  {
   v1.GET("/stocks", GetStocks)
  }

  r.Use(static.Serve("../clientLayer")) // static files have higher priority over dynamic routes
  
  r.NotFound(static.Serve("/clientLayer"))

  r.Run(":3000")

}