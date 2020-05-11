package main

import (
	"fmt"
	"github.com/cloverzrg/go-portforward/config"
	"github.com/cloverzrg/go-portforward/forwarding"
	"github.com/cloverzrg/go-portforward/logger"
	"github.com/cloverzrg/go-portforward/web"
)

// @title go-portforward
// @version 1.0

// @contact.name API Support
// @contact.url https://github.com/cloverzrg/go-portforward
// @contact.email cloverzrg@gmail.com

// @license.name go-portforward

var (
	BuildTime string
	GoVersion string
	GitHead   string
)

func main() {
	var err error
	a := forwarding.New2()
	//a.Close()
	if err != nil {
		logger.Panic(err)
	}
	err = web.Start()
	if err != nil {
		logger.Panic(err)
	}
}

func init() {
	fmt.Printf("BuildTime: %s\nGoVersion: %s\nGitHead: %s\n", BuildTime, GoVersion, GitHead)
	config.Parse("./config.json")
}
