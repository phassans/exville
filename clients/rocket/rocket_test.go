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

	testGroupName = "testGroup"
)

var (
	rClient Client
)

func newRocketChatClient(t *testing.T) {
	common.InitLogger()
	rClient = NewRocketClient(rocketURL, common.GetLogger())

	if err := rClient.InitClient(testAdminUserName, testAdminPassword); err != nil {
		panic("cannot initialize test rocket client")
	}
}

func TestClient_Login(t *testing.T) {
	newRocketChatClient(t)
}

func TestClient_CreateDeleteInfoUser(t *testing.T) {
	newRocketChatClient(t)
	{
		resp, err := rClient.CreateUser(getNewUserRequest())
		require.NoError(t, err)
		require.Equal(t, true, resp.Success)

		_, err = rClient.CreateUser(getNewUserRequest())
		require.Error(t, err)

		iresp, err := rClient.InfoUser(InfoUserRequest{testUsername})
		require.NoError(t, err)
		require.Equal(t, true, iresp.Success)

		dresp, err := rClient.DeleteUser(DeleteUserRequest{iresp.User.ID})
		require.NoError(t, err)
		require.Equal(t, true, dresp.Success)
	}
}

func TestClient_CreateDeleteGroup(t *testing.T) {
	newRocketChatClient(t)
	{
		resp, err := rClient.CreateGroup(GroupCreateRequest{testGroupName})
		require.NoError(t, err)
		require.Equal(t, true, resp.Success)

		_, err = rClient.CreateGroup(GroupCreateRequest{testGroupName})
		require.Error(t, err)

		// get Group info
		infoGroupResp, err := rClient.InfoGroup(InfoGroupRequest{testGroupName})
		require.NoError(t, err)
		require.Equal(t, true, infoGroupResp.Success)

		dresp, err := rClient.DeleteGroup(DeleteGroupRequest{infoGroupResp.Group.ID})
		require.NoError(t, err)
		require.Equal(t, true, dresp.Success)
	}
}

func TestClient_AddRemoveFromGroup(t *testing.T) {
	newRocketChatClient(t)
	{
		// create user
		createUserResp, err := rClient.CreateUser(getNewUserRequest())
		require.NoError(t, err)
		require.Equal(t, true, createUserResp.Success)

		// get user info
		infoUserResp, err := rClient.InfoUser(InfoUserRequest{testUsername})
		require.NoError(t, err)
		require.Equal(t, true, infoUserResp.Success)

		// create Group
		createGroupResp, err := rClient.CreateGroup(GroupCreateRequest{testGroupName})
		require.NoError(t, err)
		require.Equal(t, true, createGroupResp.Success)

		// get Group info
		infoGroupResp, err := rClient.InfoGroup(InfoGroupRequest{testGroupName})
		require.NoError(t, err)
		require.Equal(t, true, infoGroupResp.Success)

		// add user to Group
		addUserToGroupResp, err := rClient.AddUserToGroup(AddUserToGroupRequest{infoGroupResp.Group.ID, infoUserResp.User.ID})
		require.NoError(t, err)
		require.Equal(t, true, addUserToGroupResp.Success)

		// remove user from Group
		removeUserFromGroupResp, err := rClient.RemoveUserFromGroup(RemoveUserFromGroupRequest{infoGroupResp.Group.ID, infoUserResp.User.ID})
		require.NoError(t, err)
		require.Equal(t, true, removeUserFromGroupResp.Success)

		// delete Group
		deleteGroupResp, err := rClient.DeleteGroup(DeleteGroupRequest{infoGroupResp.Group.ID})
		require.NoError(t, err)
		require.Equal(t, true, deleteGroupResp.Success)

		// delete user
		deleteUserResp, err := rClient.DeleteUser(DeleteUserRequest{infoUserResp.User.ID})
		require.NoError(t, err)
		require.Equal(t, true, deleteUserResp.Success)
	}
}

func getNewUserRequest() CreateUserRequest {
	return CreateUserRequest{testName, testUserEmail, testUsername, testUserPassword}
}
