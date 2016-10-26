package musdk

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func (c *Client) httpReq(uri string, method string, buffer *bytes.Buffer) (string, int, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, uri, buffer)
	req.Header.Set("Token", c.token)
	res, err := client.Do(req)

	defer res.Body.Close()
	if err != nil {
		return "", 0, err
	}
	body, err := ioutil.ReadAll(res.Body)
	return string(body), res.StatusCode, err
}

func (c *Client) httpGet(uri string) (string, int, error) {
	return c.httpReq(uri, http.MethodGet, nil)
}

func (c *Client) httpPost(uri, data string) (string, int, error) {
	return c.httpReq(uri, http.MethodPost, bytes.NewBuffer([]byte(data)))
}
