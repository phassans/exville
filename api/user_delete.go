package controller

import (
	"context"
	"fmt"

	"github.com/phassans/exville/common"
	"github.com/phassans/exville/engines"
)

type (
	userDeleteRequest struct {
		UserID engines.UserID `json:"userId"`
	}

	userDeleteResponse struct {
		Request userDeleteRequest `json:"request,omitempty"`
		Error   *APIError         `json:"error,omitempty"`
	}

	userDeleteEndpoint struct{}
)

var userDelete postEndpoint = userDeleteEndpoint{}

func (r userDeleteEndpoint) Execute(ctx context.Context, rtr *router, requestI interface{}) (interface{}, error) {
	request := requestI.(userDeleteRequest)
	if err := r.Validate(requestI); err != nil {
		return userDeleteResponse{}, err
	}

	err := rtr.engines.DeleteUser(request.UserID)
	result := userDeleteResponse{Request: request, Error: NewAPIError(err)}
	return result, err
}

func (r userDeleteEndpoint) Validate(request interface{}) error {
	input := request.(userDeleteRequest)
	if input.UserID <= 0 {
		return common.ValidationError{Message: fmt.Sprint("delete user failed, missing fields")}
	}
	return nil
}

func (r userDeleteEndpoint) GetPath() string {
	return "/delete"
}

func (r userDeleteEndpoint) HTTPRequest() interface{} {
	return userDeleteRequest{}
}
