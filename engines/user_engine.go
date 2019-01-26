package engines

import (
	"fmt"

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
		Login(Username, Password) (User, error)

		/*GetUserProfile(Username)
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
	var userId UserID
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

func (u *userEngine) Login(username Username, password Password) (User, error) {
	return u.dbEngine.GetUserByUserNameAndPassword(username, password)
}

func (u *userEngine) getAndProcessUserProfile(linkedInURL LinkedInURL, userId UserID) error {
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
	if err := u.dbEngine.UpdateUserWithNameAndReference(FirstName(profile.User.Firstname), LastName(profile.User.LastName), FileName(profile.FileName), userId); err != nil {
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

func (u *userEngine) addUserToSchools(profile phantom.Profile, userID UserID) error {
	for _, school := range profile.Schools {
		schoolID, err := u.dbEngine.AddSchoolIfNotPresent(SchoolName(school.SchoolName), Degree(school.Degree), FieldOfStudy(school.FieldOfStudy))
		if err != nil {
			return err
		}

		if err := u.dbEngine.AddUserToSchool(userID, schoolID, FromYear(school.FromYear), ToYear(school.ToYear)); err != nil {
			return err
		}
	}

	return nil
}

func (u *userEngine) addUserToCompanies(profile phantom.Profile, userID UserID) error {
	for _, company := range profile.Companies {
		companyID, err := u.dbEngine.AddCompanyIfNotPresent(CompanyName(company.CompanyName), Location(company.Location))
		if err != nil {
			return err
		}

		if err := u.dbEngine.AddUserToCompany(userID, companyID, Title(company.Title), FromYear(company.FromYear), ToYear(company.ToYear)); err != nil {
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
