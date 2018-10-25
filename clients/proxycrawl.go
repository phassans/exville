package clients

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const token = "ZpEooicQ2-WsoMwF3tYDvw"

func CrawlLink(link string) {
	url := url.QueryEscape(link)
	resp, _ := http.Get(fmt.Sprintf("https://api.proxycrawl.com/?token=%s&url=%s", token, url))

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("response Body: ", string(body))
}
