package user

import (
	"fmt"

	"github.com/phassans/exville/engines/database"

	"github.com/phassans/exville/clients/phantom"
	"github.com/phassans/exville/clients/rocket"
	"github.com/rs/zerolog"
)

type (
	userEngine struct {
		rClient  rocket.Client
		pClient  phantom.Client
		dbEngine database.DatabaseEngine
		logger   zerolog.Logger
	}

	UserEngine interface {
		SignUp(phantom.Username, phantom.Password, phantom.LinkedInURL) error
		/*Login(Username, Password) error

		GetUserProfile(Username)
		GetProfileByURL(LinkedInURL) (Profile, error)

		CreateOrVerifyGroups([]Group) error
		AddUserToGroups(User, []Group)
		RemoveUserFromGroups(User, []Group)*/
	}
)

func NewUserEngine(rClient rocket.Client, pClient phantom.Client, dbEngine database.DatabaseEngine, logger zerolog.Logger) (UserEngine, error) {
	return &userEngine{
		rClient,
		pClient,
		dbEngine,
		logger,
	}, nil
}

func (u *userEngine) SignUp(username phantom.Username, password phantom.Password, linkedInURL phantom.LinkedInURL) error {
	// add user to db
	var userId database.UserID
	var err error

	// add user
	userId, err = u.dbEngine.AddUser(username, password, linkedInURL)
	if err != nil {
		return err
	}

	if err := u.getAndProcessUserProfile(linkedInURL, userId); err != nil {
		return err
	}

	return nil
}

func (u *userEngine) getAndProcessUserProfile(linkedInURL phantom.LinkedInURL, userId database.UserID) error {
	// get userProfile
	profile, err := u.pClient.GetUserProfile(string(linkedInURL))
	if err != nil {
		return err
	}

	if err := u.addUserToSchools(profile, userId); err != nil {
		return nil
	}

	if err := u.addUserToCompanies(profile, userId); err != nil {
		return nil
	}

	// update user preferences
	if err := u.dbEngine.UpdateUserWithNameAndReference(profile.User.Firstname, profile.User.LastName, profile.FileName, userId); err != nil {
		return err
	}

	// update user preferences
	grps, err := u.dbEngine.AddGroupsToUser(userId)
	if err != nil {
		return err
	}
	fmt.Println(grps)

	return nil
}

func (u *userEngine) addUserToSchools(profile phantom.Profile, userID database.UserID) error {
	for _, school := range profile.Schools {
		schoolID, err := u.dbEngine.AddSchoolIfNotPresent(school.SchoolName, school.Degree, school.FieldOfStudy)
		if err != nil {
			return err
		}

		if err := u.dbEngine.AddUserToSchool(userID, schoolID, school.FromYear, school.ToYear); err != nil {
			return err
		}
	}

	return nil
}

func (u *userEngine) addUserToCompanies(profile phantom.Profile, userID database.UserID) error {
	for _, company := range profile.Companies {
		companyID, err := u.dbEngine.AddCompanyIfNotPresent(company.CompanyName, company.Location)
		if err != nil {
			return err
		}

		if err := u.dbEngine.AddUserToCompany(userID, companyID, company.Title, company.FromYear, company.ToYear); err != nil {
			return err
		}
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
