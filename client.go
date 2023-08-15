package musdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
)

var (
	UpdateTrafficFail     = errors.New("Update Traffic Failed ")
	UpdateOnlineCountFail = errors.New("Update Online Count Failed ")
	StatusCodeError       = errors.New("Status code is not OK ")
)

type Client struct {
	baseUrl string
	nodeID  string
	token   string
	sType   int // service Type

	userTraffic map[int64]UserTrafficLog
	userTFmu    *sync.Mutex
	logger      *slog.Logger
}

func NewClient(baseUrl, token, nodeID string, sType int, logger *slog.Logger) *Client {

	return &Client{
		baseUrl:     baseUrl,
		token:       token,
		nodeID:      nodeID,
		sType:       sType,
		userTraffic: make(map[int64]UserTrafficLog),
		userTFmu:    new(sync.Mutex),
		logger:      logger,
	}
}

func ClientFromEnv(logger *slog.Logger) *Client {
	return NewClient(env("MU_URI"), env("MU_TOKEN"), env("MU_NODE_ID"), envInt("MU_SERVICE_TYPE"), logger)
}

func (c *Client) SetLogger(l *slog.Logger) {
	c.logger = l
}

func (c *Client) getUsersUri() string {
	return fmt.Sprintf("%s/nodes/%s/users", c.baseUrl, c.nodeID)
}

func (c *Client) getV2rayUsersUri() string {
	return fmt.Sprintf("%s/nodes/%s/v2rayUsers", c.baseUrl, c.nodeID)
}

func (c *Client) postTrafficUri() string {
	return fmt.Sprintf("%s/nodes/%s/traffic", c.baseUrl, c.nodeID)
}

func (c *Client) postIpUri() string {
	return fmt.Sprintf("%s/nodes/%s/ip", c.baseUrl, c.nodeID)
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
