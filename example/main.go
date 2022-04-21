package main

import (
	"log"
	"syservice"

	"github.com/Dereking/syservice"
	"github.com/gin-gonic/gin"
)

var logger service.Logger
var service *syservice.Syservice

type Program struct {
}

func (p *Program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}
	go p.run()
	return nil
}

func (p *Program) run() {
	// Do work here

	logger.Info("listen and serve on", yamlConfig.App.Port)

	r := gin.Default()
	r.GET("/", index)
	r.Run(yamlConfig.App.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (p *Program) Stop(s service.Service) error {
	logger.Info("Stopping!")
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	conf := NewServiceConfig("syservice", "syservice", "syservice")
	prg := &Program{}
	service = syservice.NewSyservice(prg, conf)
	if err != nil {
		log.Fatal(err)
	}

	service.Run()

}
