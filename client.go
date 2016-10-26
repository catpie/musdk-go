package musdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	UpdateTrafficFail     = errors.New("Update Traffic Failed ")
	UpdateOnlineCountFail = errors.New("Update Online Count Failed")
	StatusCodeError       = errors.New("Status code is not OK")
)

type Client struct {
	baseUrl string
	nodeId  int
	token   string
}

func NewClient(baseUrl, token string, nodeId int) *Client {
	client := new(Client)
	client.baseUrl = baseUrl
	client.token = token
	client.nodeId = nodeId
	return client
}

func (c *Client) getUsersUri() string {
	return fmt.Sprintf("%s/nodes/%d/users", c.baseUrl, c.nodeId)
}

func (c *Client) postTrafficUri() string {
	return fmt.Sprintf("%s/nodes/%d/traffic", c.baseUrl, c.nodeId)
}

func (c *Client) GetUsers() ([]User, error) {
	var users []User
	resp, statusCode, err := c.httpGet(c.getUsersUri())

	if err != nil {
		return users, err
	}

	if statusCode != http.StatusOK {
		return users, errors.New(fmt.Sprintf("status code: %d", statusCode))
	}

	var ret UserDataRet
	err = json.Unmarshal([]byte(resp), &ret)
	if err != nil {
		return users, err
	}

	return ret.Data, nil
}

func (c *Client) UpdateTraffic() error {
	return nil
}
