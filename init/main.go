package main

import (
	"flag"
	"github.com/3boku/Go-Server/init/cmd"
)

var configPathFlag = flag.String("confg", "config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
