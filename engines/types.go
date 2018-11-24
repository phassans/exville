package engines

type (
	FLName string

	Username string

	Password string

	LinkedInURL string

	SchoolName string

	Degree string

	FieldOfStudy string

	FromYear int32

	ToYear int32

	CompanyName string

	Title string

	Location string

	UserID int64

	SchoolID int64

	CompanyID int64

	Group string

	User struct {
		name FLName
	}

	Company struct {
		companyName CompanyName
		fromYear    FromYear
		toYear      ToYear
		title       Title
		location    Location
	}

	School struct {
		schoolName   SchoolName
		degree       Degree
		fieldOfStudy FieldOfStudy
		fromYear     FromYear
		toYear       ToYear
	}

	Profile struct {
		user      User
		companies []Company
		schools   []School
	}
)
