package engines

import (
	"github.com/phassans/exville/clients/rocket"
	"github.com/rs/zerolog"
)

type (
	userEngine struct {
		rClient     rocket.Client
		credentials rocket.AdminCredentials
		logger      zerolog.Logger
	}

	UserEngine interface {
		CreateOrCheckUserChannels(channels []Channel) error
	}
)

func NewUserEngine(client rocket.Client, userName string, password string, logger zerolog.Logger) (UserEngine, error) {
	resp, err := client.Login(rocket.UserLoginRequest{userName, password})
	if err != nil {
		return nil, err
	}

	return &userEngine{
		client,
		rocket.AdminCredentials{resp.Data.AuthToken, resp.Data.UserID},
		logger,
	}, nil
}

func (u *userEngine) CreateOrCheckUserChannels(channels []Channel) error {
	logger := u.logger
	resp, err := u.rClient.CreateChannel(rocket.ChannelCreateRequest{"channel1"}, u.credentials)
	if err != nil {
		return err
	}

	logger.Info().Msgf("response: %s", resp.Success)
	return nil
}
