package main

import(
  "btolsen131/CaliGo/config"
  "btolsen131/CaliGo/listener"
  "fmt"
)

func main(){
  cfg := config.LoadConfig("config.json")

  fmt.Println("Listen Ports:",cfg.ListenPorts)
  fmt.Println("Forward Destinations:", cfg.ForwardDest)

  listener.Listen(cfg.ListenPorts,cfg.ForwardDest)

}
