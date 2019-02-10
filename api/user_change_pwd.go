package controller

import (
	"context"
	"fmt"
	"strings"

	"github.com/phassans/exville/common"
	"github.com/phassans/exville/engines"
)

type (
	userChangePwdRequest struct {
		UserID   engines.UserID   `json:"userId"`
		Password engines.Password `json:"password,omitempty"`
	}

	userChangePwdResponse struct {
		Request userChangePwdRequest `json:"request,omitempty"`
		Error   *APIError            `json:"error,omitempty"`
	}

	userChangePwdEndpoint struct{}
)

var userChangePwd postEndpoint = userChangePwdEndpoint{}

func (r userChangePwdEndpoint) Execute(ctx context.Context, rtr *router, requestI interface{}) (interface{}, error) {
	request := requestI.(userChangePwdRequest)
	if err := r.Validate(requestI); err != nil {
		return userChangePwdResponse{}, err
	}

	err := rtr.engines.ChangePassword(request.UserID, request.Password)
	result := userChangePwdResponse{Request: request, Error: NewAPIError(err)}
	return result, err
}

func (r userChangePwdEndpoint) Validate(request interface{}) error {
	input := request.(userChangePwdRequest)
	if input.UserID <= 0 ||
		strings.TrimSpace(string(input.Password)) == "" {
		return common.ValidationError{Message: fmt.Sprint("changepwd failed, missing fields")}
	}
	return nil
}

func (r userChangePwdEndpoint) GetPath() string {
	return "/changepwd"
}

func (r userChangePwdEndpoint) HTTPRequest() interface{} {
	return userChangePwdRequest{}
}
