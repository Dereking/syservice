package syservice

import (
	"github.com/kardianos/service"
)

type Status = service.Status

type IService interface {
	Start(s *Syservice) error

	// Stop provides a place to clean up program execution before it is terminated.
	// It should not take more then a few seconds to execute.
	// Stop should not call os.Exit directly in the function.
	Stop(s *Syservice) error

	//
	// Shutdown provides a place to clean up program execution when the system is being shutdown.
	// It is essentially the same as Stop but for the case where machine is being shutdown/restarted
	// instead of just normally stopping the service. Stop won't be called when Shutdown is.
	Shutdown(s *Syservice) error
}
