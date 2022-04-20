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
	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"

	svcConfig := &service.Config{
		Name:        "testsvr",
		DisplayName: "test api server",
		Description: "this is a test api server",
		//Arguments:   []string{"start"},
		Option: options,
	}

	gin.SetMode(gin.ReleaseMode)

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
				log.Fatal("install err:", err)
			}
			fmt.Println("Service has been installed: ", svcConfig.Name)
			return
		case "uninstall", "remove":
			err = s.Uninstall()
			if err != nil {
				log.Fatal("uninstall err:", err)
			}
			fmt.Println("Service has been uninstalled.")
			return
		case "start":
			fmt.Println("Service is staring.")
			err = s.Start()
			//err = s.Run()
			if err != nil {
				log.Fatal("start err", err)
			}
			return
		case "stop":
			err = s.Stop()
			if err != nil {
				log.Fatal("stop err", err)
			}
			return
		case "restart":
			err = s.Restart()
			if err != nil {
				log.Fatal("restart err", err)
			}
			return
		}
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}

}
