package rocket

import (
	"testing"

	"github.com/phassans/exville/common"
	"github.com/stretchr/testify/require"
)

const (
	rocketURL         = "http://65.255.36.179:3000"
	testAdminUserName = "pramod"
	testAdminPassword = "123456"

	testName         = "banana"
	testUsername     = "iambanana"
	testUserEmail    = "banana@gmail.com"
	testUserPassword = "banana123"

	testChannelName = "testchannel"
)

var (
	adminCredentials AdminCredentials
	rClient          Client
)

func newRocketChatClient(t *testing.T) {
	common.InitLogger()
	rClient = NewRocketClient(rocketURL, common.GetLogger())
	loginResp, err := rClient.Login(UserLoginRequest{testAdminUserName, testAdminPassword})
	require.NoError(t, err)
	require.Equal(t, "success", loginResp.Status)
	adminCredentials = AdminCredentials{AuthToken: loginResp.Data.AuthToken, UserId: loginResp.Data.UserID}
}

func TestClient_Login(t *testing.T) {
	newRocketChatClient(t)
}

func TestClient_CreateDeleteInfoUser(t *testing.T) {
	newRocketChatClient(t)
	{
		resp, err := rClient.CreateUser(getNewUserRequest(), adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, resp.Success)

		_, err = rClient.CreateUser(getNewUserRequest(), adminCredentials)
		require.Error(t, err)

		iresp, err := rClient.InfoUser(InfoUserRequest{testUsername}, adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, iresp.Success)

		dresp, err := rClient.DeleteUser(DeleteUserRequest{iresp.User.ID}, adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, dresp.Success)
	}
}

func TestClient_CreateDeleteChannel(t *testing.T) {
	newRocketChatClient(t)
	{
		resp, err := rClient.CreateChannel(ChannelCreateRequest{testChannelName}, adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, resp.Success)

		_, err = rClient.CreateChannel(ChannelCreateRequest{testChannelName}, adminCredentials)
		require.Error(t, err)

		dresp, err := rClient.DeleteChannel(DeleteChannelRequest{testChannelName}, adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, dresp.Success)
	}
}

func TestClient_AddRemoveFromChannel(t *testing.T) {
	newRocketChatClient(t)
	{
		// create user
		resp, err := rClient.CreateUser(getNewUserRequest(), adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, resp.Success)

		// get user info
		iresp, err := rClient.InfoUser(InfoUserRequest{testUsername}, adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, iresp.Success)

		// create channel
		cresp, err := rClient.CreateChannel(ChannelCreateRequest{testChannelName}, adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, cresp.Success)

		// get channel info
		icresp, err := rClient.InfoChannel(InfoChannelRequest{testChannelName}, adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, icresp.Success)

		// add user to channel
		auresp, err := rClient.AddUserToChannel(AddUserToChannelRequest{icresp.Channel.ID, iresp.User.ID}, adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, auresp.Success)

		// remove user from channel
		remUser, err := rClient.RemoveUserFromChannel(RemoveUserFromChannelRequest{icresp.Channel.ID, iresp.User.ID}, adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, remUser.Success)

		// delete channel
		dresp, err := rClient.DeleteChannel(DeleteChannelRequest{testChannelName}, adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, dresp.Success)

		// delete user
		duresp, err := rClient.DeleteUser(DeleteUserRequest{iresp.User.ID}, adminCredentials)
		require.NoError(t, err)
		require.Equal(t, true, duresp.Success)
	}
}

func getNewUserRequest() CreateUserRequest {
	return CreateUserRequest{testName, testUserEmail, testUsername, testUserPassword}
}
