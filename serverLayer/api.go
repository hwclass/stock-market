package main

import (
  "github.com/gin-gonic/gin"
  "encoding/json"
  "fmt"
  "time"
)

type Stock struct {
  Id int64 `db:"id" json:"id"`
  Firm string `db:"firm" json:"firm"`
  Data string `db:"data" json:"data"`
} 

func GetStocks(c *gin.Context) {
  c.Writer.Header().Add("cache-control", "no-cache")
  c.Writer.Header().Add("connection", "keep-alive")
  c.Writer.Header().Add("content-type", "text/event-stream")
  type Stocks []Stock
  stocks := Stocks{
    Stock{Firm: "facebook", Data: "Oliver"},
    Stock{Firm: "apple", Data: "Malcom"},
  }
  result, err := json.Marshal(stocks)
  if err != nil {
    fmt.Println("error:", err)
  }
  c.String(200, "id: " + time.Now().Local().String() + "\n")
  c.String(200, "data: " + string(result) + "\n")
  c.Done()
}

func Cors() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
    c.Next()
  }
}
  
func main() {
  router := gin.Default()
  router.Use(Cors())
  v1 := router.Group("api/v1") 
  {
   v1.GET("/stocks", GetStocks)
  }

  router.Static("/app", "../clientLayer")

  // Listen and server on 0.0.0.0:8080
  router.Run(":3000")
}
