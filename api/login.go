package controller

import (
	"context"
	"fmt"
	"strings"

	"github.com/phassans/exville/common"
)

type (
	loginRequest struct {
		UserName string `json:"userName"`
		Password string `json:"password,omitempty"`
	}

	loginResponse struct {
		loginRequest
		Error *APIError `json:"error,omitempty"`
	}

	loginEndpoint struct{}
)

var login postEndpoint = loginEndpoint{}

func (r loginEndpoint) Execute(ctx context.Context, rtr *router, requestI interface{}) (interface{}, error) {
	request := requestI.(loginRequest)

	if err := r.Validate(requestI); err != nil {
		return nil, err
	}

	result := loginResponse{loginRequest: request, Error: NewAPIError(nil)}
	return result, nil
}

func (r loginEndpoint) Validate(request interface{}) error {
	input := request.(loginRequest)
	if strings.TrimSpace(input.UserName) == "" ||
		strings.TrimSpace(input.Password) == "" {
		return common.ValidationError{Message: fmt.Sprint("login failed, missing fields")}
	}
	return nil
}

func (r loginEndpoint) GetPath() string {
	return "/login"
}

func (r loginEndpoint) HTTPRequest() interface{} {
	return loginRequest{}
}
