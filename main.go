package main

import(
  "btolsen131/CaliGo/config"
  "btolsen131/CaliGo/receiver"
  "fmt"
)

func main(){
  cfg := config.LoadConfig("config.json")

  fmt.Println("Listen Ports:",cfg.ListenPorts)
  fmt.Println("Forward Destinations:", cfg.ForwardDest)

  receiver.Listen(cfg.ListenPorts)

}
