package configuration

// Configuration contains the application configuration settings
type Configuration struct {
	Address         string // the address to serve the application
	EnableLogin     bool   // controls enabling the login endpoint
	EnableModules   bool   // controls enabling the modules endpoint
	EnableProviders bool   // controls enabling the providers endpoint
	Endpoint        string // the endpoint which servers the Terraform registry
	Port            string // the port to serve the application
	Verbose         bool   // configures verbosity of output logs
}

// New creates and returns a new default Configuration
func New() *Configuration {
	return &Configuration{}
}
