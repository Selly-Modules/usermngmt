package config

type Configuration struct {
	EmailIsUnique       bool
	PhoneNumberIsUnique bool
}

var (
	c *Configuration
)

// Set ...
func Set(instance *Configuration) {
	c = instance
}

// GetInstance ...
func GetInstance() *Configuration {
	return c
}
