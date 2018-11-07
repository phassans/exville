package rocket

import (
	"encoding/json"
	"fmt"
)

func (c *client) CreateChannel(request ChannelCreateRequest, params AdminCredentials) (ChannelCreateResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, createChannel, params)
	if err != nil {
		var errResp ChannelErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return ChannelCreateResponse{}, err
		}
		logger = logger.With().
			Bool("success", errResp.Success).
			Str("error", errResp.Error).
			Str("errorType", errResp.ErrorType).
			Logger()
		logger.Error().Msgf("CreateChannel returned with error")
		return ChannelCreateResponse{}, fmt.Errorf("CreateChannel returned with error: %s, type: %s", errResp.Error, errResp.ErrorType)
	}

	// read response
	var resp ChannelCreateResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on ChannelCreateResponse")
		return ChannelCreateResponse{}, err
	}

	return resp, nil
}

func (c *client) DeleteChannel(request DeleteChannelRequest, params AdminCredentials) (DeleteChannelResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, deleteChannel, params)
	if err != nil {
		var errResp ChannelErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return DeleteChannelResponse{}, err
		}
		logger = logger.With().
			Bool("success", errResp.Success).
			Str("error", errResp.Error).
			Str("errorType", errResp.ErrorType).
			Logger()
		logger.Error().Msgf("DeleteChannel returned with error")
		return DeleteChannelResponse{}, fmt.Errorf("DeleteChannel returned with error: %s, type: %s", errResp.Error, errResp.ErrorType)
	}

	// read response
	var resp DeleteChannelResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on DeleteChannelResponse")
		return DeleteChannelResponse{}, err
	}

	return resp, nil
}

func (c *client) InfoChannel(request InfoChannelRequest, params AdminCredentials) (InfoChannelResponse, error) {
	logger := c.logger

	urlParams := map[string]string{"roomName": request.RoomName}

	response, err := c.DoGet(urlParams, infoChannel, params)
	if err != nil {
		var errResp ErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return InfoChannelResponse{}, err
		}

		logger = logger.With().
			Str("code", errResp.Error).
			Str("error", errResp.Message).
			Str("status", errResp.Status).
			Logger()
		logger.Error().Msgf("InfoChannel returned with error")
		return InfoChannelResponse{}, fmt.Errorf("InfoChannel returned with error: %s and code: %s", errResp.Message, errResp.Error)

	}

	// read response
	var resp InfoChannelResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on InfoChannelResponse")
		return InfoChannelResponse{}, err
	}

	return resp, nil
}
