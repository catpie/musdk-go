package musdk

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) httpReq(uri string, method string, buffer string) (string, int, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, uri, strings.NewReader(buffer))
	if err != nil {
		return "", 0, err
	}
	req.Header.Set("Token", c.token)
	req.Header.Set("ServiceType", strconv.Itoa(c.sType))
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)

	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return string(body), res.StatusCode, err
}

func (c *Client) httpGet(uri string) (string, int, error) {
	return c.httpReq(uri, http.MethodGet, "")
}

func (c *Client) httpPost(uri, data string) (string, int, error) {
	return c.httpReq(uri, http.MethodPost, data)
}
