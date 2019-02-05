package controller

import (
	"context"
	"fmt"
	"strings"

	"github.com/phassans/exville/common"
	"github.com/phassans/exville/engines"
)

type (
	signUpRequest struct {
		UserName    engines.Username    `json:"userName"`
		Password    engines.Password    `json:"password,omitempty"`
		LinkedInURL engines.LinkedInURL `json:"linkedInURL"`
	}

	signUpResponse struct {
		signUpRequest
		Error *APIError `json:"error,omitempty"`
	}

	signUpEndpoint struct{}
)

var signUp postEndpoint = signUpEndpoint{}

func (r signUpEndpoint) Execute(ctx context.Context, rtr *router, requestI interface{}) (interface{}, error) {
	request := requestI.(signUpRequest)

	if err := r.Validate(requestI); err != nil {
		return nil, err
	}

	err := rtr.engines.SignUp(request.UserName, request.Password, request.LinkedInURL)
	result := signUpResponse{signUpRequest: request, Error: NewAPIError(err)}
	return result, err
}

func (r signUpEndpoint) Validate(request interface{}) error {
	input := request.(signUpRequest)
	if strings.TrimSpace(string(input.UserName)) == "" ||
		strings.TrimSpace(string(input.Password)) == "" {
		return common.ValidationError{Message: fmt.Sprint("signUp failed, missing fields")}
	}
	return nil
}

func (r signUpEndpoint) GetPath() string {
	return "/signup"
}

func (r signUpEndpoint) HTTPRequest() interface{} {
	return signUpRequest{}
}
