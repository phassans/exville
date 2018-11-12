package phantom

import (
	"fmt"
	"testing"

	"github.com/phassans/exville/common"
	"github.com/stretchr/testify/require"
)

const (
	phantomURL = "https://phantombuster.com"
)

var (
	pClient Client
)

func newPhantomClient(t *testing.T) {
	common.InitLogger()
	pClient = NewPhantomClient(phantomURL, common.GetLogger())
}

func TestClient_CrawlUrl(t *testing.T) {
	newPhantomClient(t)
	{

		resp, err := pClient.CrawlUrl("https://www.linkedin.com/in/pramod-shashidhara-21568923")
		require.NoError(t, err)
		for _, obj := range resp.Data.ResultObject {
			for _, jobs := range obj.Jobs {
				fmt.Println(jobs.CompanyName)
				fmt.Println(jobs.Location)
			}

			for _, school := range obj.Schools {
				fmt.Println(school.SchoolName)
				fmt.Println(school.DateRange)
				fmt.Println(school.Degree)
				fmt.Println(school.DegreeSpec)
			}
		}
	}
}
