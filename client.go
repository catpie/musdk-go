package musdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/weeon/contract"
)

var (
	UpdateTrafficFail     = errors.New("Update Traffic Failed ")
	UpdateOnlineCountFail = errors.New("Update Online Count Failed ")
	StatusCodeError       = errors.New("Status code is not OK ")
)

type Client struct {
	baseUrl string
	nodeId  int
	token   string
	sType   int // service Type

	userTraffic map[int64]UserTrafficLog
	userTFmu    *sync.Mutex
	logger      contract.Logger
}

func NewClient(baseUrl, token string, nodeId, sType int, logger contract.Logger) *Client {

	return &Client{
		baseUrl:     baseUrl,
		token:       token,
		nodeId:      nodeId,
		sType:       sType,
		userTraffic: make(map[int64]UserTrafficLog),
		userTFmu:    new(sync.Mutex),
		logger:      logger,
	}
}

func ClientFromEnv(logger contract.Logger) *Client {
	return NewClient(env("MU_URI"), env("MU_TOKEN"), envInt("MU_NODE_ID"), envInt("MU_SERVICE_TYPE"), logger)
}

func (c *Client) SetLogger(l contract.Logger) {
	c.logger = l
}

func (c *Client) getUsersUri() string {
	return fmt.Sprintf("%s/nodes/%d/users", c.baseUrl, c.nodeId)
}

func (c *Client) getV2rayUsersUri() string {
	return fmt.Sprintf("%s/nodes/%d/v2rayUsers", c.baseUrl, c.nodeId)
}

func (c *Client) postTrafficUri() string {
	return fmt.Sprintf("%s/nodes/%d/traffic", c.baseUrl, c.nodeId)
}

func (c *Client) postIpUri() string {
	return fmt.Sprintf("%s/nodes/%d/ip", c.baseUrl, c.nodeId)
}

func (c *Client) getNodesUri() string {
	return fmt.Sprintf("%s/nodes", c.baseUrl)
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

func (c *Client) UpdateTraffic(logs []UserTrafficLog) error {
	data, err := json.Marshal(logs)
	if err != nil {
		return err
	}
	_, statusCode, err := c.httpPost(c.postTrafficUri(), string(data))
	if err != nil {
		return err
	}
	if statusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("status code: %d", statusCode))
	}
	return nil
}

func (c *Client) GetNodes() ([]Node, error) {
	resp, statusCode, err := c.httpGet(c.getNodesUri())
	if err != nil {
		return nil, err
	}
	if statusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("status code: %d", statusCode))
	}
	var ret NodeDataRet
	err = json.Unmarshal([]byte(resp), &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}

type ipReq struct {
	IP string `json:"ip"`
}

func (c *Client) PostIP(ip string) error {
	var req = ipReq{
		IP: ip,
	}
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	_, statusCode, err := c.httpPost(c.postIpUri(), string(data))
	if err != nil {
		return err
	}
	if statusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("status code: %d", statusCode))
	}
	return nil
}
