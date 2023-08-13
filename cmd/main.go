package main

import (
	"elasticSearch/cmd/apps"
	"elasticSearch/config"
	"flag"
)

var configFlag = flag.String("config", "./config.toml", "configuration toml file path")

func main() {
	flag.Parse()
	cfg := config.NewConfig(*configFlag)
	app := apps.NewApps(cfg)
	app.Wait()
}
