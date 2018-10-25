package rocket

import (
	"fmt"
	"testing"
)

func TestCreateChannel(t *testing.T) {
	req := UserLoginRequest{Username: "pramod", Password: "123456"}
	resp, err := Login(req)
	if err != nil {
		fmt.Println(err)
	}

	CreateChannel(ChannelCreateRequest{"channel1"}, OtherRequestParams{resp.Data.AuthToken, resp.Data.UserID})
}
