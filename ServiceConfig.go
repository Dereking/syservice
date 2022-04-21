package syservice

type ServiceConfig struct {
	Name        string
	DisplayName string
	Description string
	Arguments   []string

	Restart           string //	options["Restart"] = "on-success"
	SuccessExitStatus string //options["SuccessExitStatus"] = "1 2 8 SIGKILL"

	LogOutput  string // options["LogOutput"] = "/var/log/service.log"
	User       string // options["User"] = "root"
	ConfigFile string // options["ConfigFile"] = "/etc/service.conf"
}

func NewServiceConfig(name, dispalyName, Description string) *ServiceConfig {
	return &ServiceConfig{
		Name:              name,
		DisplayName:       dispalyName,
		Description:       Description,
		Restart:           "on-success",
		SuccessExitStatus: "1 2 8 SIGKILL",
		LogOutput:         "/var/log/service.log",
		User:              "root",
		ConfigFile:        "/etc/service.conf",
	}
}
