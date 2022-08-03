package syservice

import (
	"github.com/gin-gonic/gin"
)

type Program struct {
	userService ISerivce
}

func NewProgram(userService_ ISerivce) *Program {
	return &Program{
		userService: userService_,
	}
}

func (p *Program) Start(s *Syservice) error {

	go p.userService.run()
	return nil
}

// Stop provides a place to clean up program execution before it is terminated.
// It should not take more then a few seconds to execute.
// Stop should not call os.Exit directly in the function.

func (p *Program) Stop(s *Syservice) error {
	return nil
}

//
// Shutdown provides a place to clean up program execution when the system is being shutdown.
// It is essentially the same as Stop but for the case where machine is being shutdown/restarted
// instead of just normally stopping the service. Stop won't be called when Shutdown is.

func (p *Program) Shutdown(s *syservice.Syservice) error {
	return nil
}

func (p *Program) run() {
	// Do work here

	//logger.Info("listen and serve on", yamlConfig.App.Port)

	r := gin.Default()
	r.GET("/", index)
	r.Run(yamlConfig.App.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
