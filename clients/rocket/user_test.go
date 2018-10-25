package rocket

import "testing"

func TestLogin(t *testing.T) {
	req := UserLoginRequest{Username: "pramod", Password: "123456"}
	Login(req)
}
