package cmd

import (
	"fmt"
	"github.com/3boku/Go-Server/config"
	"github.com/3boku/Go-Server/network"
	"github.com/3boku/Go-Server/repository"
	"github.com/3boku/Go-Server/service"
)

type Cmd struct {
	config     *config.Config
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filepath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)
	fmt.Println(c.config.Server.Port)
	return c
}
