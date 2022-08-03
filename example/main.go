package main

import (
	"log"

	"github.com/Dereking/syservice"
	"github.com/gin-gonic/gin"
)

var service *syservice.Syservice

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
