package main

import (
	"github.com/tonet-me/tonet-core/cmd"
	"github.com/tonet-me/tonet-core/config"
)

func main() {
	cfg := config.C()

	cmd.StartServe(cfg)
}
