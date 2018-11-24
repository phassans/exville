package engines

import "github.com/phassans/exville/clients/phantom"

const (
	testDatabaseHost     = "localhost"
	testDataPort         = "5432"
	testDatabase         = "viraagh"
	testDatabaseUsername = "pshashidhara"
	testDatabasePassword = "banana123"

	testUserFirstName phantom.FirstName   = "Pramod"
	testUserLastName  phantom.LastName    = "HS"
	testUserName      phantom.Username    = "Viraagh"
	testPassword      phantom.Password    = "123456"
	testLinkedInURL   phantom.LinkedInURL = "https://linkedin.com/viraagh"

	testSchool       phantom.SchoolName   = "Colorado State University"
	testDegree       phantom.Degree       = "Masters"
	testFieldOfStudy phantom.FieldOfStudy = "Computer Science"

	testCompany  phantom.CompanyName = "Hungry Hour"
	testLocation phantom.Location    = "Sunnyvale"
	testTitle    phantom.Title       = "Developer"

	testFromYear phantom.FromYear = 2017
	testToYear   phantom.ToYear   = 2018

	testFileName phantom.FileName = "pramod.shashidhara.json"
)
