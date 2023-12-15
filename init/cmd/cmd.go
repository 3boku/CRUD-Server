package cmd

import (
	"fmt"
	"github.com/3boku/Go-Server/config"
	"github.com/3boku/Go-Server/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filepath),
		network: network.NewNetwork(),
	}

	network.NewNetwork().ServerStart(c.config.Server.Port)
	fmt.Println(c.config.Server.Port)
	return c
}
