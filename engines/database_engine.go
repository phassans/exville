package engines

import (
	"database/sql"
	"time"

	"github.com/phassans/exville/helper"
	"github.com/rs/zerolog"
)

type (
	databaseEngine struct {
		sql    *sql.DB
		logger zerolog.Logger
	}

	DatabaseEngine interface {
		AddUser(flName FLName, username Username, password Password, linkedInURL LinkedInURL) (UserID, error)
		DeleteUser(username Username) error
		UpdateUser(flName FLName, username Username, password Password, linkedInURL LinkedInURL) error

		AddSchool(school School, degree Degree, fieldOfStudy FieldOfStudy) (SchoolID, error)
		DeleteSchool(school School, degree Degree, fieldOfStudy FieldOfStudy) error

		AddCompany(company Company, location Location) (CompanyID, error)
		DeleteCompany(company Company, location Location) error

		AddUserToSchool(userID UserID, schoolID SchoolID, fromYear FromYear, toYear ToYear) error
		RemoveUserFromSchool(userID UserID, schoolID SchoolID) error

		AddUserToCompany(userID UserID, companyID CompanyID, title Title, fromYear FromYear, toYear ToYear) error
		RemoveUserFromCompany(userID UserID, companyID CompanyID) error
	}
)

// NewDatabaseEngine returns an instance of userEngine
func NewDatabaseEngine(psql *sql.DB, logger zerolog.Logger) DatabaseEngine {
	return &databaseEngine{psql, logger}
}

func (d *databaseEngine) AddUser(flName FLName, username Username, password Password, linkedInURL LinkedInURL) (UserID, error) {
	var userID UserID
	err := d.sql.QueryRow("INSERT INTO viraagh_user(fl_name,username,password,linkedIn_URL,insert_time) "+
		"VALUES($1,$2,$3,$4,$5) returning user_id;",
		flName, username, password, linkedInURL, time.Now()).Scan(&userID)
	if err != nil {
		return 0, helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully added a user with ID: %d", userID)

	return userID, nil
}

func (d *databaseEngine) DeleteUser(username Username) error {
	_, err := d.sql.Exec("DELETE FROM viraagh_user WHERE username=$1", username)
	if err != nil {
		return helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully delete user: %s", username)
	return nil
}

func (d *databaseEngine) UpdateUser(flName FLName, username Username, password Password, linkedInURL LinkedInURL) error {
	return nil
}

func (d *databaseEngine) AddSchool(school School, degree Degree, fieldOfStudy FieldOfStudy) (SchoolID, error) {
	var schoolID SchoolID
	err := d.sql.QueryRow("INSERT INTO school(school,degree,field_of_study,insert_time) "+
		"VALUES($1,$2,$3,$4) returning school_id;",
		school, degree, fieldOfStudy, time.Now()).Scan(&schoolID)
	if err != nil {
		return 0, helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully added a school with ID: %d", schoolID)

	return schoolID, nil
}

func (d *databaseEngine) DeleteSchool(school School, degree Degree, fieldOfStudy FieldOfStudy) error {
	_, err := d.sql.Exec("DELETE FROM school WHERE school=$1 AND degree = $2 AND field_of_study = $3", school, degree, fieldOfStudy)
	if err != nil {
		return helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully delete school: %s", school)
	return nil
}

func (d *databaseEngine) AddCompany(company Company, location Location) (CompanyID, error) {
	var companyID CompanyID
	err := d.sql.QueryRow("INSERT INTO company(company,location,insert_time) "+
		"VALUES($1,$2,$3) returning company_id;",
		company, location, time.Now()).Scan(&companyID)
	if err != nil {
		return 0, helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully added a company with ID: %d", companyID)

	return companyID, nil
}

func (d *databaseEngine) DeleteCompany(company Company, location Location) error {
	_, err := d.sql.Exec("DELETE FROM company WHERE company=$1 AND location=$2", company, location)
	if err != nil {
		return helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully delete company: %s", company)
	return nil
}

func (d *databaseEngine) AddUserToSchool(userID UserID, schoolID SchoolID, fromYear FromYear, toYear ToYear) error {
	_, err := d.sql.Exec("INSERT INTO user_to_school(user_id,school_id,from_year,to_year,insert_time) "+
		"VALUES($1,$2,$3,$4,$5)",
		userID, schoolID, fromYear, toYear, time.Now())
	if err != nil {
		return helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully added a user: %d to school: %d", userID, schoolID)

	return nil
}

func (d *databaseEngine) RemoveUserFromSchool(userID UserID, schoolID SchoolID) error {
	_, err := d.sql.Exec("DELETE FROM user_to_school WHERE user_id=$1 AND school_id=$2", userID, schoolID)
	if err != nil {
		return helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully removed user: %d from school: %d", userID, schoolID)
	return nil
}

func (d *databaseEngine) AddUserToCompany(userID UserID, companyID CompanyID, title Title, fromYear FromYear, toYear ToYear) error {
	_, err := d.sql.Exec("INSERT INTO user_to_company(user_id,company_id,title,from_year,to_year,insert_time) "+
		"VALUES($1,$2,$3,$4,$5,$6)",
		userID, companyID, title, fromYear, toYear, time.Now())
	if err != nil {
		return helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully added a user: %d to company: %d", userID, companyID)

	return nil
}

func (d *databaseEngine) RemoveUserFromCompany(userID UserID, companyID CompanyID) error {
	_, err := d.sql.Exec("DELETE FROM user_to_company WHERE user_id=$1 AND company_id=$2", userID, companyID)
	if err != nil {
		return helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully removed user: %d from company: %d", userID, companyID)
	return nil
}
