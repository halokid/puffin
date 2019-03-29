package main

import (
  "context"
  "github.com/micro/go-micro"
  ipdbpro "github.com/r00tjimmy/puffin/ipdb/srv/proto/ipdbpro"
  "log"
  "time"
)

type IpDb struct {

}

func (ipDb *IpDb) GetIpInfo(ctx context.Context, req *ipdbpro.Request, rsp *ipdbpro.Response) error {
  log.Println("recv IpDb.GetIpInfo req")
  rsp.IpInfo = "ip query " + req.Ip
  return nil
}

func main() {
  service := micro.NewService(
    micro.Name("go.micro.srv.puffin"),
    micro.RegisterTTL(time.Second * 30),        // fixme: timeout set??
    micro.RegisterInterval(time.Second * 10),                // what's mean??
  )

  service.Init()

  _ = ipdbpro.RegisterIpDbHandler(service.Server(), new(IpDb))

  if err := service.Run(); err != nil {
    log.Fatal(err)
  }
}


