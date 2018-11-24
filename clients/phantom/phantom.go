package phantom

import "github.com/rs/zerolog"

type (
	client struct {
		baseURL string
		logger  zerolog.Logger
	}

	Client interface {
		CrawlUrl(string) (CrawlResponse, error)

		GetUserProfile(string) (Profile, error)
		SaveUserProfile(resp CrawlResponse) (string, error)

		GetSchoolsFromResponse(resp CrawlResponse) ([]School, error)
		GetCompaniesFromResponse(resp CrawlResponse) ([]Company, error)
		GetUserFromResponse(resp CrawlResponse) User
	}
)

// NewPhantomClient returns a new phantom client
func NewPhantomClient(baseURL string, logger zerolog.Logger) Client {
	return &client{baseURL, logger}
}
