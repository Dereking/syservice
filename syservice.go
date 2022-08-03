package syservice

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/kardianos/service"
)

type Syservice struct {
	logger      service.Logger
	baseService *service.Service
	program     *Program
}

func NewSyservice(pserv IService, conf ServiceConfig) (*Syservice, error) {
	svcConfig := &service.Config{
		Name:        conf.Name,
		DisplayName: conf.DisplayName,
		Description: conf.Description,
	}

	baseProgr := NewProgram(pserv)

	baseService, err := service.New(baseProgr, svcConfig)
	if err != nil {
		return nil, err
	}

	log_, err = s.Logger(nil)
	if err != nil {
		return nil, err
	}

	svc := &Syservice{
		baseService: baseService,
		logger:      log_,
		program:     baseProgr,
	}
	//svc.Init()
	return svc, nil
}

func (this *Syservice) Run() {

	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)

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
