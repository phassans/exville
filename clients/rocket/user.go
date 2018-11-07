package rocket

import (
	"encoding/json"
	"fmt"
)

func (c *client) Login(request UserLoginRequest) (UserLoginResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, login, AdminCredentials{})
	if err != nil {
		var errResp ErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return UserLoginResponse{}, err
		}

		logger = logger.With().
			Str("code", errResp.Error).
			Str("error", errResp.Message).
			Str("status", errResp.Status).
			Logger()
		logger.Error().Msgf("login returned with error")
		return UserLoginResponse{}, fmt.Errorf("login returned with error: %s and code: %s", errResp.Message, errResp.Error)
	}

	// read response
	var resp UserLoginResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on UserLoginResponse")
		return UserLoginResponse{}, err
	}

	return resp, nil
}

func (c *client) CreateUser(request CreateUserRequest, params AdminCredentials) (CreateUserResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, createUser, params)
	if err != nil {
		var errResp ErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return CreateUserResponse{}, err
		}

		logger = logger.With().
			Str("code", errResp.Error).
			Str("error", errResp.Message).
			Str("status", errResp.Status).
			Logger()
		logger.Error().Msgf("create user returned with error")
		return CreateUserResponse{}, fmt.Errorf("CreateUser returned with error: %s and code: %s", errResp.Message, errResp.Error)

	}

	// read response
	var resp CreateUserResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on CreateUserResponse")
		return CreateUserResponse{}, err
	}

	return resp, nil
}

func (c *client) DeleteUser(request DeleteUserRequest, params AdminCredentials) (DeleteUserResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request, deleteUser, params)
	if err != nil {
		var errResp ErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return DeleteUserResponse{}, err
		}

		logger = logger.With().
			Str("code", errResp.Error).
			Str("error", errResp.Message).
			Str("status", errResp.Status).
			Logger()
		logger.Error().Msgf("delete user returned with error")
		return DeleteUserResponse{}, fmt.Errorf("DeleteUser returned with error: %s and code: %s", errResp.Message, errResp.Error)

	}

	// read response
	var resp DeleteUserResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on DeleteUserResponse")
		return DeleteUserResponse{}, err
	}

	return resp, nil
}

func (c *client) InfoUser(request InfoUserRequest, params AdminCredentials) (InfoUserResponse, error) {
	logger := c.logger

	urlParams := map[string]string{"username": request.Username}

	response, err := c.DoGet(urlParams, infoUser, params)
	if err != nil {
		var errResp ErrorResponse
		err = json.Unmarshal(response, &errResp)
		if err != nil {
			logger = logger.With().Str("error", err.Error()).Logger()
			logger.Error().Msgf("unmarshal error on ErrorResponse")
			return InfoUserResponse{}, err
		}

		logger = logger.With().
			Str("code", errResp.Error).
			Str("error", errResp.Message).
			Str("status", errResp.Status).
			Logger()
		logger.Error().Msgf("info user returned with error")
		return InfoUserResponse{}, fmt.Errorf("InfoUser returned with error: %s and code: %s", errResp.Message, errResp.Error)

	}

	// read response
	var resp InfoUserResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on InfoUserResponse")
		return InfoUserResponse{}, err
	}

	return resp, nil
}
