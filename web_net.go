package musdk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) genGetUsersUrl() string {
	return fmt.Sprintf("%s/users?key=%s", c.baseUrl, c.key)
}

func (c *Client) genUserTrafficUrl(id int) string {
	return fmt.Sprintf("%s/users/%d/traffic?key=%s", c.baseUrl, id, c.key)
}

func (c *Client) genNodeOnlineCountUrl(id int) string {
	return fmt.Sprintf("%s/nodes/%d/online_count?key=%s", c.baseUrl, id, c.key)
}

func (c *Client) genNodeInfoUrl(id int) string {
	return fmt.Sprintf("%s/nodes/%d/info?key=%s", c.baseUrl, id, c.key)
}

func (c *Client) httpGet(urlStr string) (string, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (c *Client) httpPostUserTraffic(userId int, u, d string) (string, error) {
	nodeId := strconv.Itoa(c.nodeId)
	urlStr := c.genUserTrafficUrl(userId)
	resp, err := http.PostForm(urlStr,
		url.Values{"u": {u}, "d": {d}, "node_id": {nodeId}})

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (c *Client) httpPostNodeOnlineCount(count int) (string, error) {
	urlStr := c.genNodeOnlineCountUrl(c.nodeId)
	countStr := strconv.Itoa(count)
	resp, err := http.PostForm(urlStr,
		url.Values{"count": {countStr}})

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (c *Client) httpPostNodeInfo(load, uptime string) (string, error) {
	urlStr := c.genNodeInfoUrl(c.nodeId)
	resp, err := http.PostForm(urlStr,
		url.Values{"load": {load}, "uptime": {uptime}})

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
