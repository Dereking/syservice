package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
)

var logger service.Logger

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
	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"

	svcConfig := &service.Config{
		Name:        "twapisvr",
		DisplayName: "twitter api server",
		Description: "this is a twitter api server",

		//deamon 3
		//dependencies : []string{"dummy.service"}
		// dependencies: []string{
		// 	"Requires=network.target",
		// 	"After=network-online.target syslog.target"},
		Option: options,
	}

	//gin.SetMode(gin.ReleaseMode)

	prg := &Program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			err = s.Install()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Service has been installed.")
			return
		case "uninstall":
		case "remove":
			err = s.Uninstall()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Service has been uninstalled.")
			return
		case "start":
			err = s.Start()
			if err != nil {
				log.Fatal(err)
			}
			return
		case "stop":
			err = s.Stop()
			if err != nil {
				log.Fatal(err)
			}
			return
		case "restart":
			err = s.Restart()
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}

}
