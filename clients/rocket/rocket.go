package rocket

import "github.com/rs/zerolog"

type (
	client struct {
		baseURL string
		logger  zerolog.Logger
	}

	Client interface {
		// baseMethod for making a post call
		DoPost(interface{}, string, AdminCredentials) ([]byte, error)

		// user login, mostly to obtain accessToken & userId
		Login(UserLoginRequest) (UserLoginResponse, error)
		InfoUser(InfoUserRequest, AdminCredentials) (InfoUserResponse, error)
		CreateUser(CreateUserRequest, AdminCredentials) (CreateUserResponse, error)
		DeleteUser(DeleteUserRequest, AdminCredentials) (DeleteUserResponse, error)

		// creates a new channel
		CreateChannel(ChannelCreateRequest, AdminCredentials) (ChannelCreateResponse, error)
		DeleteChannel(DeleteChannelRequest, AdminCredentials) (DeleteChannelResponse, error)
		InfoChannel(InfoChannelRequest, AdminCredentials) (InfoChannelResponse, error)

		AddUserToChannel(AddUserToChannelRequest, AdminCredentials) (AddUserToChannelResponse, error)
		RemoveUserFromChannel(RemoveUserFromChannelRequest, AdminCredentials) (RemoveUserFromChannelResponse, error)
	}
)

// NewRocketClient returns a new rocket client
func NewRocketClient(baseURL string, logger zerolog.Logger) Client {
	return &client{baseURL, logger}
}
