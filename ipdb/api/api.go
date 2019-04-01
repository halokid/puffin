package main

import (
  "context"
  "encoding/json"
  "github.com/micro/go-micro"
  "github.com/micro/go-micro/errors"
  ipdbpro "github.com/r00tjimmy/puffin/ipdb/srv/proto/ipdbpro"
  "log"
  api "github.com/micro/micro/api/proto"
  "strings"
)

type IpDb struct {
  Client  ipdbpro.IpDbService
}

func (ipDb *IpDb) GetIpInfo(ctx context.Context, req *api.Request, rsp *api.Response) error {
  log.Print("Recv IpDb.GetIpInfo API request")

  ip, ok := req.Get["Ip"]
  if !ok || len(ip.Values) == 0 {
    return errors.BadRequest("go.micro.api.puffin.ipdb", "Ip can not be blank")
  }

  response, err := ipDb.Client.GetIpInfo(ctx, &ipdbpro.Request{
    Ip:  strings.Join(ip.Values, " "),
  })
  if err != nil {
    return err
  }

  rsp.StatusCode = 200
  b, _ := json.Marshal(map[string]string{
    "msg":  response.IpInfo,
  })
  rsp.Body = string(b)

  return nil
}

func main() {
  service := micro.NewService(
    micro.Name("go.micro.api.puffin.ipdb"),
  )

  service.Init()

  _ = service.Server().Handle(
    service.Server().NewHandler(
      &IpDb{Client: ipdbpro.NewIpDbService("go.micro.srv.puffin.ipdb", service.Client())},
    ),
  )

  if err := service.Run(); err != nil {
    log.Fatal(err)
  }
}





