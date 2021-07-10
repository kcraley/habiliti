package configuration

// Configuration contains the application configuration settings
type Configuration struct {
	Verbose bool
}

// New creates and returns a new default Configuration
func New() *Configuration {
	return &Configuration{}
}
