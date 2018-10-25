package rocket

import "github.com/rs/zerolog"

type (
	client struct {
		baseURL string
		logger  zerolog.Logger
	}

	Client interface {
		// baseMethod for making a post call
		DoPost(request interface{}, requestType string, params AdminCredentials) ([]byte, error)

		// user login, mostly to obtain accessToken & userId
		Login(request UserLoginRequest) (UserLoginResponse, error)

		// creates a new channel
		CreateChannel(request ChannelCreateRequest, params AdminCredentials) (ChannelCreateResponse, error)
	}
)

// NewRocketClient returns a new rocket client
func NewRocketClient(baseURL string, logger zerolog.Logger) Client {
	return &client{baseURL, logger}
}
