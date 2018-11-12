package engines

import (
	"testing"

	"github.com/phassans/exville/common"
	"github.com/phassans/exville/db"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

var (
	roach  db.Roach
	logger zerolog.Logger
	engine DatabaseEngine
)

func newDataBaseEngine(t *testing.T) {
	logger = common.GetLogger()
	if logger.Log() == nil {
		{
			common.InitLogger()
		}
	}

	var err error
	if roach.Db == nil {
		roach, err = db.New(db.Config{Host: testDatabaseHost, Port: testDataPort, User: testDatabaseUsername, Password: testDatabasePassword, Database: testDatabase})
		if err != nil {
			logger.Fatal().Msgf("could not connect to db. errpr %s", err)
		}
	}

	engine = NewDatabaseEngine(roach.Db, logger)
	require.NotEmpty(t, engine)
}

func TestDatabaseEngine_AddDeleteUser(t *testing.T) {
	newDataBaseEngine(t)
	{
		userID, err := engine.AddUser(testUser, testUserName, testPassword, testLinkedInURL)
		require.NoError(t, err)
		require.NotEmpty(t, userID)

		err = engine.DeleteUser(testUserName)
		require.NoError(t, err)
	}
}

func TestDatabaseEngine_AddDeleteSchool(t *testing.T) {
	newDataBaseEngine(t)
	{
		schoolID, err := engine.AddSchool(testSchool, testDegree, testFieldOfStudy)
		require.NoError(t, err)
		require.NotEmpty(t, schoolID)

		err = engine.DeleteSchool(testSchool, testDegree, testFieldOfStudy)
		require.NoError(t, err)
	}
}

func TestDatabaseEngine_AddDeleteCompany(t *testing.T) {
	newDataBaseEngine(t)
	{
		companyID, err := engine.AddCompany(testCompany, testLocation)
		require.NoError(t, err)
		require.NotEmpty(t, companyID)

		err = engine.DeleteCompany(testCompany, testLocation)
		require.NoError(t, err)
	}
}

func TestDatabaseEngine_AddRemoveUserSchool(t *testing.T) {
	newDataBaseEngine(t)
	{
		userID, err := engine.AddUser(testUser, testUserName, testPassword, testLinkedInURL)
		require.NoError(t, err)
		require.NotEmpty(t, userID)

		schoolID, err := engine.AddSchool(testSchool, testDegree, testFieldOfStudy)
		require.NoError(t, err)
		require.NotEmpty(t, schoolID)

		err = engine.AddUserToSchool(userID, schoolID, testFromYear, testToYear)
		require.NoError(t, err)

		err = engine.RemoveUserFromSchool(userID, schoolID)
		require.NoError(t, err)

		err = engine.DeleteSchool(testSchool, testDegree, testFieldOfStudy)
		require.NoError(t, err)

		err = engine.DeleteUser(testUserName)
		require.NoError(t, err)
	}
}

func TestDatabaseEngine_AddRemoveUserCompany(t *testing.T) {
	newDataBaseEngine(t)
	{
		userID, err := engine.AddUser(testUser, testUserName, testPassword, testLinkedInURL)
		require.NoError(t, err)
		require.NotEmpty(t, userID)

		companyID, err := engine.AddCompany(testCompany, testLocation)
		require.NoError(t, err)
		require.NotEmpty(t, companyID)

		err = engine.AddUserToCompany(userID, companyID, testTitle, testFromYear, testToYear)
		require.NoError(t, err)

		err = engine.RemoveUserFromCompany(userID, companyID)
		require.NoError(t, err)

		err = engine.DeleteCompany(testCompany, testLocation)
		require.NoError(t, err)

		err = engine.DeleteUser(testUserName)
		require.NoError(t, err)
	}
}
