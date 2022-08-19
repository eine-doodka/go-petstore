package main

import (
	"context"
	"example.com/prj/internal/app/apiserver"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "cfg file path")
}

func main() {
	flag.Parse()
	cfg := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, &cfg)
	if err != nil {
		panic(err)
	}
	s := apiserver.NewServer(cfg)
	ctx := context.Background()
	if err := s.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
