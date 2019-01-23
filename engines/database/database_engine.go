package database

import (
	"database/sql"
	"fmt"
	"strings"
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
		// User Methods
		AddUser(username phantom.Username, password phantom.Password, linkedInURL phantom.LinkedInURL) (UserID, error)
		DeleteUser(username phantom.Username) error
		UpdateUserWithNameAndReference(name phantom.FirstName, lastName phantom.LastName, fileName phantom.FileName, id UserID) error

		// School Methods
		AddSchoolIfNotPresent(school phantom.SchoolName, degree phantom.Degree, fieldOfStudy phantom.FieldOfStudy) (SchoolID, error)
		DeleteSchool(school phantom.SchoolName, degree phantom.Degree, fieldOfStudy phantom.FieldOfStudy) error
		GetSchoolID(phantom.SchoolName, phantom.Degree, phantom.FieldOfStudy) (SchoolID, error)

		// Company Methods
		AddCompanyIfNotPresent(company phantom.CompanyName, location phantom.Location) (CompanyID, error)
		DeleteCompany(company phantom.CompanyName, location phantom.Location) error
		GetCompanyID(phantom.CompanyName, phantom.Location) (CompanyID, error)

		// UserToSchool
		AddUserToSchool(userID UserID, schoolID SchoolID, fromYear phantom.FromYear, toYear phantom.ToYear) error
		RemoveUserFromSchool(userID UserID, schoolID SchoolID) error

		// UserToCompany
		AddUserToCompany(userID UserID, companyID CompanyID, title phantom.Title, fromYear phantom.FromYear, toYear phantom.ToYear) error
		RemoveUserFromCompany(userID UserID, companyID CompanyID) error

		AddGroupsToUser(userID UserID) ([]phantom.Group, error)
		GetGroupsByUserID(userID UserID) ([]phantom.Group, error)
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

func (d *databaseEngine) GetSchoolsByUserID(userID UserID) ([]phantom.School, error) {
	rows, err := d.sql.Query("SELECT school.school_name, school.degree, school.field_of_study,user_to_school.from_year, user_to_school.to_year "+
		"FROM school INNER JOIN user_to_school ON school.school_id = user_to_school.school_id "+
		"WHERE user_to_school.user_id=$1", userID)
	if err != nil {
		return nil, helper.DatabaseError{DBError: err.Error()}
	}

	defer rows.Close()

	var schools []phantom.School
	for rows.Next() {
		var school phantom.School
		err = rows.Scan(&school.SchoolName, &school.Degree, &school.FieldOfStudy, &school.FromYear, &school.ToYear)
		if err != nil {
			return nil, helper.DatabaseError{DBError: err.Error()}
		}
		schools = append(schools, school)
	}

	return schools, nil
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

func (d *databaseEngine) GetCompaniesByUserID(userID UserID) ([]phantom.Company, error) {
	rows, err := d.sql.Query("SELECT company.company_name, company.location "+
		"FROM company INNER JOIN user_to_company ON company.company_id = user_to_company.company_id "+
		"WHERE user_to_company.user_id=$1", userID)
	if err != nil {
		return nil, helper.DatabaseError{DBError: err.Error()}
	}

	defer rows.Close()

	var companies []phantom.Company
	for rows.Next() {
		var company phantom.Company
		err = rows.Scan(&company.CompanyName, &company.Location)
		if err != nil {
			return nil, helper.DatabaseError{DBError: err.Error()}
		}
		companies = append(companies, company)
	}

	return companies, nil
}

func (d *databaseEngine) AddGroupsToUser(userID UserID) ([]phantom.Group, error) {
	groups, err := d.getGroupsBySchoolsAndCompanies(userID)
	if err != nil {
		return nil, err
	}

	// Note: only add unique groups
	grpMap := make(map[phantom.Group]bool)
	var uniqGroups []phantom.Group
	for _, group := range groups {
		if !grpMap[group] {
			// insert into school
			_, err = d.sql.Exec("INSERT INTO user_to_groups(user_id,group_name,status) VALUES($1,$2,$3);", userID, group, true)
			if err != nil {
				return nil, helper.DatabaseError{DBError: err.Error()}
			}
			grpMap[group] = true
			uniqGroups = append(uniqGroups, group)

			d.logger.Info().Msgf("user with ID:%d joined group: %s", userID, group)
		}
	}

	return uniqGroups, nil
}

func (d *databaseEngine) GetGroupsByUserID(userID UserID) ([]phantom.Group, error) {
	rows, err := d.sql.Query("SELECT group_name FROM user_to_groups "+
		"WHERE user_id=$1 AND status=$2", userID, true)
	if err != nil {
		return nil, helper.DatabaseError{DBError: err.Error()}
	}

	defer rows.Close()

	var groups []phantom.Group

	for rows.Next() {
		var group phantom.Group
		err = rows.Scan(&group)
		if err != nil {
			return nil, helper.DatabaseError{DBError: err.Error()}
		}
		groups = append(groups, group)

	}

	return groups, nil
}

func (d *databaseEngine) getGroupsBySchoolsAndCompanies(userID UserID) ([]phantom.Group, error) {
	var groups []phantom.Group
	groupsSchools, err := d.GetSchoolsByUserID(userID)
	if err != nil {
		return nil, err
	}
	groups = append(groups, d.schoolsToGroups(groupsSchools)...)

	groupsCompanies, err := d.GetCompaniesByUserID(userID)
	if err != nil {
		return nil, err
	}
	groups = append(groups, d.companiesToGroups(groupsCompanies)...)
	return groups, nil
}

func (d *databaseEngine) schoolsToGroups(schools []phantom.School) []phantom.Group {
	var grps []phantom.Group

	for _, school := range schools {
		// add schoolName
		schoolName := strings.Replace(string(school.SchoolName), " ", "", -1)
		grps = append(grps, phantom.Group(schoolName))

		degree := strings.Replace(string(school.Degree), " ", "", -1)
		fieldOfStudy := strings.Replace(string(school.FieldOfStudy), " ", "", -1)

		// add combination of school, degree & fieldOfStudy
		groupName := fmt.Sprintf("%s-%s-%s-%d-%d", schoolName, degree, fieldOfStudy, school.FromYear, school.ToYear)
		grps = append(grps, phantom.Group(groupName))

	}
	return grps
}

func (d *databaseEngine) companiesToGroups(companies []phantom.Company) []phantom.Group {
	var grps []phantom.Group

	for _, company := range companies {
		// add companyName
		companyName := strings.Replace(string(company.CompanyName), " ", "", -1)
		grps = append(grps, phantom.Group(companyName))

		location := strings.Replace(string(company.Location), " ", "", -1)
		// add combination of companyName & location
		groupName := fmt.Sprintf("%s-%s", companyName, location)
		grps = append(grps, phantom.Group(groupName))

	}
	return grps
}
