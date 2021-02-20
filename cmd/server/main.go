package main

import (
	"flag"
	"go-seed/internal/api"
	"go-seed/pkg/config"
)

func main() {
	cfgPath := flag.String("c", "./cmd/server/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}