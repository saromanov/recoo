package recoo

// Config defines configuration
type Config struct {
	Build Build `yaml:"build"`
}

// Build defined build stage
type Build struct {
	Image     string `yaml:"image"`
	Entryfile string `yaml:"entryfile"`
}
