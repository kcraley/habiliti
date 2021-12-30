package configuration

// Configuration contains the application configuration settings
type Configuration struct {
	Address string // the address to serve the application
	Port    string // the port to serve the application
	Verbose bool   // configures verbosity of output logs
}

// New creates and returns a new default Configuration
func New() *Configuration {
	return &Configuration{}
}
