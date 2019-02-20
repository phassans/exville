package cloudinary

import (
	"io"

	"github.com/rs/zerolog"
)

type (
	client struct {
		logger zerolog.Logger
	}

	Client interface {
		Upload(values map[string]io.Reader) error
	}
)

// NewCloudinaryClient returns a new cloudinary client
func NewCloudinaryClient(logger zerolog.Logger) Client {
	return &client{logger}
}
