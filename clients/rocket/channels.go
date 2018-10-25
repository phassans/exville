package rocket

import (
	"encoding/json"
	"fmt"
	"time"
)

type (
	ChannelCreateRequest struct {
		Name string `json:"name"`
	}

	ChannelCreateResponse struct {
		Channel struct {
			ID         string `json:"_id"`
			Name       string `json:"name"`
			Fname      string `json:"fname"`
			T          string `json:"t"`
			Msgs       int    `json:"msgs"`
			UsersCount int    `json:"usersCount"`
			U          struct {
				ID       string `json:"_id"`
				Username string `json:"username"`
			} `json:"u"`
			CustomFields struct {
			} `json:"customFields"`
			Ts        time.Time   `json:"ts"`
			Ro        bool        `json:"ro"`
			SysMes    bool        `json:"sysMes"`
			Default   bool        `json:"default"`
			UpdatedAt time.Time   `json:"_updatedAt"`
			Lm        interface{} `json:"lm"`
		} `json:"channel"`
		Success bool `json:"success"`
	}

	ChannelErrorResponse struct {
		Success   bool   `json:"success"`
		Error     string `json:"error"`
		ErrorType string `json:"errorType"`
	}
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
		return ChannelCreateResponse{}, fmt.Errorf("CreateChannel returned with error: %s, type: %d", errResp.Error, errResp.ErrorType)
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
