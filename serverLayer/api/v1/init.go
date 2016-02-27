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

/*
func EventStream() gin.HandlerFunc {
  return func(c *gin.Context) {
    
  }
}
*/

func Cors() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
    c.Next()
  }
}

func main() {

  r := gin.Default()

  r.Use(Cors())

  v1 := r.Group("api/v1")
  {
   v1.GET("/stocks", GetStocks)
  }

  /*
  res.writeHead(200, {
    'content-type': 'text/event-stream',
    'cache-control': 'no-cache',
    'connection': 'keep-alive'
  });
  console.dir(JSON.stringify(founders));
  res.write('id: ' + (new Date()).toLocaleTimeString() + '\n');
  res.write("data: " + JSON.stringify(founders) + '\n\n');
  */

  r.Run(":8080")

}