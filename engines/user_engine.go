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
		SignUp(phantom.Username, phantom.Password, phantom.LinkedInURL) error
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

func (u *userEngine) SignUp(username phantom.Username, password phantom.Password, linkedInURL phantom.LinkedInURL) error {
	// add user to db
	var userId UserID
	var err error

	// add user
	userId, err = u.dbEngine.AddUser(username, password, linkedInURL)
	if err != nil {
		return err
	}

	// get userProfile
	profile, err := u.pClient.GetUserProfile(string(linkedInURL))
	if err != nil {
		return err
	}

	// add schools
	var schoolIDs []SchoolID
	for _, school := range profile.Schools {
		id, err := u.dbEngine.AddSchoolIfNotPresent(school.SchoolName, school.Degree, school.FieldOfStudy)
		if err != nil {
			return err
		}

		schoolIDs = append(schoolIDs, id)
	}

	// add companies
	var companyIDs []CompanyID
	for _, company := range profile.Companies {
		id, err := u.dbEngine.AddCompanyIfNotPresent(company.CompanyName, company.Location)
		if err != nil {
			return err
		}

		companyIDs = append(companyIDs, id)
	}

	// add user to school
	for index, school := range profile.Schools {
		if err := u.dbEngine.AddUserToSchool(userId, schoolIDs[index], school.FromYear, school.ToYear); err != nil {
			return err
		}
	}

	// add user to company
	for index, company := range profile.Companies {
		if err := u.dbEngine.AddUserToCompany(userId, companyIDs[index], company.Title, company.FromYear, company.ToYear); err != nil {
			return err
		}
	}

	// update user preferences
	if err := u.dbEngine.UpdateUserWithNameAndReference(profile.User.Firstname, profile.User.LastName, profile.FileName, userId); err != nil {
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
