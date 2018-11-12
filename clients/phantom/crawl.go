package phantom

import (
	"encoding/json"
	"fmt"
)

func (c *client) CrawlUrl(linkedInURL string) (CrawlResponse, error) {
	arg := Argument{
		SessionCookie: sessionCookie,
		ProfileUrls:   []string{linkedInURL},
		NoDatabase:    true,
	}

	return c.doCrawlUrl(
		CrawlRequest{Output: output,
			Argument: arg,
		},
	)
}

func (c *client) doCrawlUrl(request CrawlRequest) (CrawlResponse, error) {
	logger := c.logger

	response, err := c.DoPost(request)
	if err != nil {
		logger.Error().Msgf("CrawlUrl returned with error")
		return CrawlResponse{}, fmt.Errorf("CrawlUrl returned with error: %s", err)
	}

	// read response
	var resp CrawlResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		logger = logger.With().Str("error", err.Error()).Logger()
		logger.Error().Msgf("unmarshal error on CrawlResponse")
		return CrawlResponse{}, err
	}

	return resp, nil
}
