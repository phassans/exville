package engines

import (
	"database/sql"
	"time"

	"github.com/phassans/exville/clients/phantom"
	"github.com/phassans/exville/helper"
	"github.com/rs/zerolog"
)

type (
	databaseEngine struct {
		sql    *sql.DB
		logger zerolog.Logger
	}

	DatabaseEngine interface {
		AddUser(username phantom.Username, password phantom.Password, linkedInURL phantom.LinkedInURL) (UserID, error)
		DeleteUser(username phantom.Username) error
		UpdateUserWithNameAndReference(name phantom.FirstName, lastName phantom.LastName, fileName phantom.FileName, id UserID) error

		AddSchoolIfNotPresent(school phantom.SchoolName, degree phantom.Degree, fieldOfStudy phantom.FieldOfStudy) (SchoolID, error)
		DeleteSchool(school phantom.SchoolName, degree phantom.Degree, fieldOfStudy phantom.FieldOfStudy) error
		GetSchoolID(phantom.SchoolName, phantom.Degree, phantom.FieldOfStudy) (SchoolID, error)

		AddCompanyIfNotPresent(company phantom.CompanyName, location phantom.Location) (CompanyID, error)
		DeleteCompany(company phantom.CompanyName, location phantom.Location) error
		GetCompanyID(phantom.CompanyName, phantom.Location) (CompanyID, error)

		AddUserToSchool(userID UserID, schoolID SchoolID, fromYear phantom.FromYear, toYear phantom.ToYear) error
		RemoveUserFromSchool(userID UserID, schoolID SchoolID) error

		AddUserToCompany(userID UserID, companyID CompanyID, title phantom.Title, fromYear phantom.FromYear, toYear phantom.ToYear) error
		RemoveUserFromCompany(userID UserID, companyID CompanyID) error
	}
)

// NewDatabaseEngine returns an instance of userEngine
func NewDatabaseEngine(psql *sql.DB, logger zerolog.Logger) DatabaseEngine {
	return &databaseEngine{psql, logger}
}

func (d *databaseEngine) AddUser(username phantom.Username, password phantom.Password, linkedInURL phantom.LinkedInURL) (UserID, error) {
	var userID UserID
	err := d.sql.QueryRow("INSERT INTO viraagh_user(username,password,linkedIn_URL,insert_time) "+
		"VALUES($1,$2,$3,$4) returning user_id;", username, password, linkedInURL, time.Now()).Scan(&userID)
	if err != nil {
		return 0, helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully added a user with ID: %d", userID)

	return userID, nil
}

func (d *databaseEngine) DeleteUser(username phantom.Username) error {
	_, err := d.sql.Exec("DELETE FROM viraagh_user WHERE username=$1", username)
	if err != nil {
		return helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully delete user: %s", username)
	return nil
}

func (d *databaseEngine) UpdateUserWithNameAndReference(firstName phantom.FirstName, lastName phantom.LastName, fileName phantom.FileName, id UserID) error {
	updateUserWithNameAndReferenceSQL := `UPDATE viraagh_user SET first_name = $1, last_name = $2, filename = $3 WHERE user_id=$4;`

	_, err := d.sql.Exec(updateUserWithNameAndReferenceSQL, firstName, lastName, fileName, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *databaseEngine) AddSchoolIfNotPresent(schoolName phantom.SchoolName, degree phantom.Degree, fieldOfStudy phantom.FieldOfStudy) (SchoolID, error) {
	var schoolID SchoolID
	var err error

	// check if school is present
	schoolID, err = d.GetSchoolID(schoolName, degree, fieldOfStudy)
	if err != nil {
		return 0, helper.DatabaseError{DBError: err.Error()}
	}

	if schoolID != 0 {
		d.logger.Info().Msgf("school:%s is added with ID: %d", schoolName, schoolID)
		return schoolID, nil
	}

	// insert into school
	err = d.sql.QueryRow("INSERT INTO school(school_name,degree,field_of_study,insert_time) "+
		"VALUES($1,$2,$3,$4) returning school_id;",
		schoolName, degree, fieldOfStudy, time.Now()).Scan(&schoolID)
	if err != nil {
		return 0, helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully added a school:%s with ID: %d", schoolName, schoolID)

	return schoolID, nil
}

func (d *databaseEngine) DeleteSchool(schoolName phantom.SchoolName, degree phantom.Degree, fieldOfStudy phantom.FieldOfStudy) error {
	_, err := d.sql.Exec("DELETE FROM school WHERE school_name=$1 AND degree = $2 AND field_of_study = $3", schoolName, degree, fieldOfStudy)
	if err != nil {
		return helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully delete school: %s", schoolName)
	return nil
}

func (d *databaseEngine) GetSchoolID(schoolName phantom.SchoolName, degree phantom.Degree, fieldOfStudy phantom.FieldOfStudy) (SchoolID, error) {
	var id SchoolID
	rows := d.sql.QueryRow("SELECT school_id FROM school where school_name = $1  AND degree = $2 AND field_of_study = $3;", schoolName, degree, fieldOfStudy)

	err := rows.Scan(&id)

	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, helper.DatabaseError{DBError: err.Error()}
	}

	return id, nil
}

func (d *databaseEngine) AddCompanyIfNotPresent(companyName phantom.CompanyName, location phantom.Location) (CompanyID, error) {
	var companyID CompanyID
	var err error

	// check if company is present
	companyID, err = d.GetCompanyID(companyName, location)
	if err != nil {
		return 0, helper.DatabaseError{DBError: err.Error()}
	}

	if companyID != 0 {
		d.logger.Info().Msgf("company:%s is added with ID: %d", companyName, companyID)
		return companyID, nil
	}

	err = d.sql.QueryRow("INSERT INTO company(company_name,location,insert_time) "+
		"VALUES($1,$2,$3) returning company_id;",
		companyName, location, time.Now()).Scan(&companyID)
	if err != nil {
		return 0, helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully added a company with ID: %d", companyID)

	return companyID, nil
}

func (d *databaseEngine) DeleteCompany(companyName phantom.CompanyName, location phantom.Location) error {
	_, err := d.sql.Exec("DELETE FROM company WHERE company_name=$1 AND location=$2", companyName, location)
	if err != nil {
		return helper.DatabaseError{DBError: err.Error()}
	}

	d.logger.Info().Msgf("successfully delete company: %s", companyName)
	return nil
}

func (d *databaseEngine) GetCompanyID(companyName phantom.CompanyName, location phantom.Location) (CompanyID, error) {
	var id CompanyID
	rows := d.sql.QueryRow("SELECT company_id FROM company where company_name = $1  AND location = $2;", companyName, location)

	err := rows.Scan(&id)

	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, helper.DatabaseError{DBError: err.Error()}
	}

	return id, nil
}

func (d *databaseEngine) AddUserToSchool(userID UserID, schoolID SchoolID, fromYear phantom.FromYear, toYear phantom.ToYear) error {
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

func (d *databaseEngine) AddUserToCompany(userID UserID, companyID CompanyID, title phantom.Title, fromYear phantom.FromYear, toYear phantom.ToYear) error {
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
