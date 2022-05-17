package k3s

import "github.com/saromanov/recoo/internal/config"

type K3S struct {
	cfg config.Deploy
}

// New provides initialization of k3s module
func New(cfg config.Deploy)*K3S {
	return &K3S{

	}
}