package musdk

import (
	"net/http"
	"io/ioutil"
	"bytes"
)

func httpReq(uri string, method string, buffer *bytes.Buffer) (string, int, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, uri, buffer)
	req.Header.Set("Token", GetClient().token)
	res, err := client.Do(req)

	defer res.Body.Close()
	if err != nil {
		return "", 0, err
	}
	body, err := ioutil.ReadAll(res.Body)
	return string(body), res.StatusCode, err
}

func httpGet(uri string) (string, int, error) {
	return httpReq(uri, http.MethodGet, nil)
}

func httpPost(uri, data string) (string, int, error) {
	return httpReq(uri, http.MethodPost, bytes.NewBuffer([]byte(data)))
}
