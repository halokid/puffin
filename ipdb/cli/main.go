package main

import (
  "context"
  "fmt"
  "github.com/micro/go-micro"
  ipdbpro "github.com/r00tjimmy/puffin/ipdb/srv/proto/ipdbpro"
)

func main()  {
  service := micro.NewService()
  service.Init()
  cl := ipdbpro.NewIpDbService("go.micro.srv.puffin.ipdb", service.Client())
  rsp, err := cl.GetIpInfo(context.Background(), &ipdbpro.Request{Ip: "8.8.8.8"})
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(rsp.IpInfo)
}
