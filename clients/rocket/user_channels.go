package rocket

import (
	"encoding/json"
	"fmt"
)

func (c *client) AddUserToChannel(request AddUserToChannelRequest, params AdminCredentials) (AddUserToChannelResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, addUserToChannel, params)
	if err != nil {
		var errResp ChannelErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return AddUserToChannelResponse{}, err
		}
		logger = logger.With().
			Bool("success", errResp.Success).
			Str("error", errResp.Error).
			Str("errorType", errResp.ErrorType).
			Logger()
		logger.Error().Msgf("AddUserToChannel returned with error")
		return AddUserToChannelResponse{}, fmt.Errorf("AddUserToChannel returned with error: %s, type: %s", errResp.Error, errResp.ErrorType)
	}

	// read response
	var resp AddUserToChannelResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on AddUserToChannelResponse")
		return AddUserToChannelResponse{}, err
	}

	return resp, nil
}

func (c *client) RemoveUserFromChannel(request RemoveUserFromChannelRequest, params AdminCredentials) (RemoveUserFromChannelResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, removeUserFromChannel, params)
	if err != nil {
		var errResp ChannelErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return RemoveUserFromChannelResponse{}, err
		}
		logger = logger.With().
			Bool("success", errResp.Success).
			Str("error", errResp.Error).
			Str("errorType", errResp.ErrorType).
			Logger()
		logger.Error().Msgf("RemoveUserFromChannelResponse returned with error")
		return RemoveUserFromChannelResponse{}, fmt.Errorf("RemoveUserToChannel returned with error: %s, type: %s", errResp.Error, errResp.ErrorType)
	}

	// read response
	var resp RemoveUserFromChannelResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on RemoveUserFromChannelResponse")
		return RemoveUserFromChannelResponse{}, err
	}

	return resp, nil
}
