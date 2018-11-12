package rocket

import (
	"encoding/json"
	"fmt"
)

func (c *client) AddUserToGroup(request AddUserToGroupRequest, params AdminCredentials) (AddUserToGroupResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, addUserToGroup, params)
	if err != nil {
		var errResp GroupErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return AddUserToGroupResponse{}, err
		}
		logger = logger.With().
			Bool("success", errResp.Success).
			Str("error", errResp.Error).
			Str("errorType", errResp.ErrorType).
			Logger()
		logger.Error().Msgf("AddUserToGroup returned with error")
		return AddUserToGroupResponse{}, fmt.Errorf("AddUserToGroup returned with error: %s, type: %s", errResp.Error, errResp.ErrorType)
	}

	// read response
	var resp AddUserToGroupResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on AddUserToGroupResponse")
		return AddUserToGroupResponse{}, err
	}

	return resp, nil
}

func (c *client) RemoveUserFromGroup(request RemoveUserFromGroupRequest, params AdminCredentials) (RemoveUserFromGroupResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, removeUserFromGroup, params)
	if err != nil {
		var errResp GroupErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return RemoveUserFromGroupResponse{}, err
		}
		logger = logger.With().
			Bool("success", errResp.Success).
			Str("error", errResp.Error).
			Str("errorType", errResp.ErrorType).
			Logger()
		logger.Error().Msgf("RemoveUserFromGroupResponse returned with error")
		return RemoveUserFromGroupResponse{}, fmt.Errorf("RemoveUserToGroup returned with error: %s, type: %s", errResp.Error, errResp.ErrorType)
	}

	// read response
	var resp RemoveUserFromGroupResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on RemoveUserFromGroupResponse")
		return RemoveUserFromGroupResponse{}, err
	}

	return resp, nil
}
