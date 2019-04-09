package main

import (
  ipdbpro "../srv/proto/ipdbpro"
  "context"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/micro/go-micro"
  "net/http"
)



func mainOld() {
  fmt.Println("strart.....")
  service := micro.NewService()
  service.Init()
  cl := ipdbpro.NewIpDbService("go.micro.srv.puffin.ipdb", service.Client())
  rsp, err := cl.GetIpInfo(context.Background(), &ipdbpro.Request{Ip: "8.8.8.8"})
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(rsp.IpInfo)

  // gin server
  router := gin.Default()
  /**
  router.GET("/ipdb/:ip", func(c *gin.Context) {
    ip := c.Param("ip")
    c.String(http.StatusOK, "ip info: %s", ip)
    //fmt.Println(ip)
  })
  **/

  router.GET("/ipinfo", func(c *gin.Context) {
    c.String(http.StatusOK, "ip info: %s", rsp.IpInfo)
  })
  _ = router.Run(":8089")
}


func getRpc() string {
  fmt.Println("getRpc...")
  service := micro.NewService()
  service.Init()
  cl := ipdbpro.NewIpDbService("go.micro.srv.puffin.ipdb", service.Client())
  rsp, err := cl.GetIpInfo(context.Background(), &ipdbpro.Request{Ip: "8.8.8.8"})
  if err != nil {
    fmt.Println(err)
    return "error..."
  }
  fmt.Println(rsp.IpInfo)
  return rsp.IpInfo
}


func main() {
  fmt.Println("strart.....")
  ipInfo := getRpc()
  // gin server
  router := gin.Default()
  router.GET("/ipinfo", func(c *gin.Context) {
    c.String(http.StatusOK, "ip info: %s\n", ipInfo)    // 缓存取得RPC返回的值
    c.String(http.StatusOK, "ip info: %s\n", getRpc())  // 实时取得RPC返回的值
  })
  _ = router.Run(":8089")
}


