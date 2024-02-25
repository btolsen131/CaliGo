package config

import (
  "encoding/json"
  "io/ioutil"
  "log"
)

type Config struct {
  ListenPorts []int
  ForwardDest []string
}

func LoadConfig(filePath string) *Config {
  file, err := ioutil.ReadFile(filePath)
  if err != nil {
    log.Fatalf("Error reading config file: %v", err)
  }

  var cfg Config
  err = json.Unmarshal(file, &cfg)
  if err != nil {
    log.Fatalf("Error parsing config: %v", err)
  }

  return &cfg
}
