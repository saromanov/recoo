package release

import "github.com/saromanov/recoo/internal/config"

// Release defines struct for release
type Release struct {
	cfg *config.Release
}

// New creates release stage
func New(cfg *config.Config) *Release {
	return &Release{
		cfg: cfg,
	}
}
