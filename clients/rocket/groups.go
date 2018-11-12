package rocket

import (
	"encoding/json"
	"fmt"
)

func (c *client) CreateGroup(request GroupCreateRequest, params AdminCredentials) (GroupCreateResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, createGroup, params)
	if err != nil {
		var errResp GroupErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return GroupCreateResponse{}, err
		}
		logger = logger.With().
			Bool("success", errResp.Success).
			Str("error", errResp.Error).
			Str("errorType", errResp.ErrorType).
			Logger()
		logger.Error().Msgf("CreateGroup returned with error")
		return GroupCreateResponse{}, fmt.Errorf("CreateGroup returned with error: %s, type: %s", errResp.Error, errResp.ErrorType)
	}

	// read response
	var resp GroupCreateResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on GroupCreateResponse")
		return GroupCreateResponse{}, err
	}

	return resp, nil
}

func (c *client) DeleteGroup(request DeleteGroupRequest, params AdminCredentials) (DeleteGroupResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, deleteGroup, params)
	if err != nil {
		var errResp GroupErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return DeleteGroupResponse{}, err
		}
		logger = logger.With().
			Bool("success", errResp.Success).
			Str("error", errResp.Error).
			Str("errorType", errResp.ErrorType).
			Logger()
		logger.Error().Msgf("DeleteGroup returned with error")
		return DeleteGroupResponse{}, fmt.Errorf("DeleteGroup returned with error: %s, type: %s", errResp.Error, errResp.ErrorType)
	}

	// read response
	var resp DeleteGroupResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on DeleteGroupResponse")
		return DeleteGroupResponse{}, err
	}

	return resp, nil
}

func (c *client) InfoGroup(request InfoGroupRequest, params AdminCredentials) (InfoGroupResponse, error) {
	logger := c.logger

	urlParams := map[string]string{"roomName": request.RoomName}

	response, err := c.DoGet(urlParams, infoGroup, params)
	if err != nil {
		var errResp ErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return InfoGroupResponse{}, err
		}

		logger = logger.With().
			Str("code", errResp.Error).
			Str("error", errResp.Message).
			Str("status", errResp.Status).
			Logger()
		logger.Error().Msgf("InfoGroup returned with error")
		return InfoGroupResponse{}, fmt.Errorf("InfoGroup returned with error: %s and code: %s", errResp.Message, errResp.Error)

	}

	// read response
	var resp InfoGroupResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on InfoGroupResponse")
		return InfoGroupResponse{}, err
	}

	return resp, nil
}

func (c *client) SetTypeGroup(request SetTypeGroupRequest, params AdminCredentials) (SetTypeGroupResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, setTypeGroup, params)
	if err != nil {
		var errResp GroupErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return SetTypeGroupResponse{}, err
		}
		logger = logger.With().
			Bool("success", errResp.Success).
			Str("error", errResp.Error).
			Str("errorType", errResp.ErrorType).
			Logger()
		logger.Error().Msgf("SetTypeGroup returned with error")
		return SetTypeGroupResponse{}, fmt.Errorf("SetTypeGroup returned with error: %s, type: %s", errResp.Error, errResp.ErrorType)
	}

	// read response
	var resp SetTypeGroupResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on SetTypeGroupResponse")
		return SetTypeGroupResponse{}, err
	}

	return resp, nil
}
