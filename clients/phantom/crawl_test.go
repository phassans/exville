package phantom

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/phassans/exville/common"
	"github.com/stretchr/testify/require"
)

const (
	phantomURL = "https://phantombuster.com"
	jsonFile   = "response.json"
)

var (
	pClient Client
)

func newPhantomClient(t *testing.T) {
	common.InitLogger()
	pClient = NewPhantomClient(phantomURL, common.GetLogger())
}

func jsonFileToResponseObject(fileName string) (Response, error) {
	plan, err := ioutil.ReadFile(fileName)
	if err != nil {
		return Response{}, nil
	}

	var resp Response
	err = json.Unmarshal(plan, &resp)
	if err != nil {
		return Response{}, nil
	}
	return resp, nil
}

func TestClient_CrawlUrl(t *testing.T) {
	newPhantomClient(t)
	{
		resp, err := pClient.CrawlUrl("https://www.linkedin.com/in/pramod-shashidhara-21568923")
		require.NoError(t, err)
		require.NotNil(t, resp)
	}
}

func TestClient_GetUserFromResponse(t *testing.T) {
	newPhantomClient(t)
	{
		resp, err := jsonFileToResponseObject(jsonFile)
		require.NoError(t, err)
		user := pClient.GetUserFromResponse(CrawlResponse{Data: resp})
		require.NoError(t, err)
		require.Equal(t, FirstName("Pramod"), user.Firstname)
		require.Equal(t, LastName("Shashidhara"), user.LastName)
	}
}

func TestClient_GetSchoolsFromResponse(t *testing.T) {
	newPhantomClient(t)
	{
		resp, err := jsonFileToResponseObject(jsonFile)
		require.NoError(t, err)
		schools, err := pClient.GetSchoolsFromResponse(CrawlResponse{Data: resp})
		require.NoError(t, err)
		require.Equal(t, 2, len(schools))
	}
}

func TestClient_GetCompaniesFromResponse(t *testing.T) {
	newPhantomClient(t)
	{
		resp, err := jsonFileToResponseObject(jsonFile)
		require.NoError(t, err)
		companies, err := pClient.GetCompaniesFromResponse(CrawlResponse{Data: resp})
		require.NoError(t, err)
		require.Equal(t, 12, len(companies))
	}
}

func TestClient_GetUserProfile(t *testing.T) {
	newPhantomClient(t)
	{
		profile, err := pClient.GetUserProfile("https://www.linkedin.com/in/pramod-shashidhara-21568923")
		require.NoError(t, err)
		require.NotNil(t, profile)
	}
}

func TestClient_SaveResponse(t *testing.T) {
	newPhantomClient(t)
	{
		resp, err := jsonFileToResponseObject(jsonFile)
		require.NoError(t, err)
		filename, err := pClient.SaveUserProfile(CrawlResponse{Data: resp})
		require.NoError(t, err)
		require.NotNil(t, filename)
	}
}
