package controller

import (
	"context"
	"fmt"
	"strings"

	"github.com/phassans/exville/common"
)

type (
	signUpRequest struct {
		UserName    string `json:"userName"`
		Password    string `json:"password,omitempty"`
		LinkedInURL string `json:"linkedInURL"`
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

	result := signUpResponse{signUpRequest: request, Error: NewAPIError(nil)}
	return result, nil
}

func (r signUpEndpoint) Validate(request interface{}) error {
	input := request.(signUpRequest)
	if strings.TrimSpace(input.UserName) == "" ||
		strings.TrimSpace(input.Password) == "" {
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
