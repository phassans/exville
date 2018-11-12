package phantom

import "github.com/rs/zerolog"

type (
	client struct {
		baseURL string
		logger  zerolog.Logger
	}

	Client interface {
		CrawlUrl(string) (CrawlResponse, error)
	}
)

// NewPhantomClient returns a new phantom client
func NewPhantomClient(baseURL string, logger zerolog.Logger) Client {
	return &client{baseURL, logger}
}
