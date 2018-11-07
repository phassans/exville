package rocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *client) DoPost(request interface{}, requestType string, params AdminCredentials) ([]byte, error) {
	logger := c.logger
	requestJson, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/%s", c.baseURL, apiPath, requestType)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")

	if requestType != login {
		req.Header.Set("X-Auth-Token", params.AuthToken)
		req.Header.Set("X-User-Id", params.UserId)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	logger = logger.With().Str("url", url).Str("status", resp.Status).Logger()

	if resp.StatusCode != 200 {
		logger = logger.With().Str("body", string(body)).Logger()
		logger.Error().Msgf("doPost non 200 response")
		return body, fmt.Errorf("post returned with errorCode: %d", resp.StatusCode)
	}
	logger.Info().Msgf("doPost success!")

	return body, nil
}

func (c *client) DoGet(requestParams map[string]string, requestType string, params AdminCredentials) ([]byte, error) {
	logger := c.logger

	url := fmt.Sprintf("%s/%s/%s", c.baseURL, apiPath, requestType)
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("X-Auth-Token", params.AuthToken)
	req.Header.Set("X-User-Id", params.UserId)

	q := req.URL.Query()
	for key, value := range requestParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	logger = logger.With().Str("url", url).Str("status", resp.Status).Logger()

	if resp.StatusCode != 200 {
		logger = logger.With().Str("body", string(body)).Logger()
		logger.Error().Msgf("doGet non 200 response")
		return body, fmt.Errorf("get returned with errorCode: %d", resp.StatusCode)
	}
	logger.Info().Msgf("doGet success!")

	return body, nil
}
