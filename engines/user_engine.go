package engines

import (
	"github.com/phassans/exville/clients/phantom"
	"github.com/phassans/exville/clients/rocket"
	"github.com/rs/zerolog"
)

type (
	userEngine struct {
		rClient  rocket.Client
		pClient  phantom.Client
		dbEngine DatabaseEngine
		logger   zerolog.Logger
	}

	UserEngine interface {
		SignUp(Username, Password, LinkedInURL) error
		/*Login(Username, Password) error

		GetUserProfile(Username)
		GetProfileByURL(LinkedInURL) (Profile, error)

		CreateOrVerifyGroups([]Group) error
		AddUserToGroups(User, []Group)
		RemoveUserFromGroups(User, []Group)*/
	}
)

func NewUserEngine(rClient rocket.Client, pClient phantom.Client, dbEngine DatabaseEngine, logger zerolog.Logger) (UserEngine, error) {
	return &userEngine{
		rClient,
		pClient,
		dbEngine,
		logger,
	}, nil
}

func (u *userEngine) SignUp(username Username, password Password, linkedInURL LinkedInURL) error {
	// add user to db
	_, err := u.dbEngine.AddUser("", username, password, linkedInURL)
	if err != nil {
		return err
	}

	return nil
}

/*func (u *userEngine) CreateOrCheckUserGroups(groups []Group) error {
	logger := u.logger
	resp, err := u.rClient.CreateGroup(rocket.GroupCreateRequest{"channel1"})
	if err != nil {
		return err
	}

	logger.Info().Msgf("response.json: %s", resp.Success)
	return nil
}*/
