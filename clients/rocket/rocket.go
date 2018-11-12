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

		// creates a new Group
		CreateGroup(GroupCreateRequest, AdminCredentials) (GroupCreateResponse, error)
		DeleteGroup(DeleteGroupRequest, AdminCredentials) (DeleteGroupResponse, error)
		InfoGroup(InfoGroupRequest, AdminCredentials) (InfoGroupResponse, error)
		SetTypeGroup(SetTypeGroupRequest, AdminCredentials) (SetTypeGroupResponse, error)

		AddUserToGroup(AddUserToGroupRequest, AdminCredentials) (AddUserToGroupResponse, error)
		RemoveUserFromGroup(RemoveUserFromGroupRequest, AdminCredentials) (RemoveUserFromGroupResponse, error)
	}
)

// NewRocketClient returns a new rocket client
func NewRocketClient(baseURL string, logger zerolog.Logger) Client {
	return &client{baseURL, logger}
}
