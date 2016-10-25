package musdk

import (
	"encoding/json"
	"errors"

	"github.com/golang/glog"
	"github.com/orvice/shadowsocks-go/mu/log"
	"github.com/orvice/shadowsocks-go/mu/system"
	"github.com/orvice/shadowsocks-go/mu/user"
)

var (
	client                = new(Client)
	UpdateTrafficFail     = errors.New("Update Traffic Failed ")
	UpdateOnlineCountFail = errors.New("Update Online Count Failed")
)

type Client struct {
	baseUrl string
	key     string
	nodeId  int
}

func NewClient() *Client {
	return new(Client)
}

func SetClient(c *Client) {
	client = c
}

func GetClient() *Client {
	return client
}

func (c *Client) setBaseUrl(baseUrl string) {
	c.baseUrl = baseUrl
}

func (c *Client) setKey(key string) {
	c.key = key
}

func (c *Client) setNodeId(id int) {
	c.nodeId = id
}

func (c *Client) GetUsers() ([]user.User, error) {
	var tempUser []user.User
	glog.Infoln("request to remote server:", c.genGetUsersUrl())
	res, err := c.httpGet(c.genGetUsersUrl())
	if err != nil {
		return tempUser, err
	}
	log.Log.Debug(res)
	var resData UserDataRet
	err = json.Unmarshal([]byte(res), &resData)
	if err != nil {
		return tempUser, err
	}
	userData := resData.Data
	log.Log.Info(len(userData))
	users := make([]user.User, len(userData))
	for k, v := range userData {
		users[k] = v
	}
	return users, nil
}

func (c *Client) UpdateTraffic(userId int, u, d string) error {
	res, err := c.httpPostUserTraffic(userId, u, d)
	if err != nil {
		return nil
	}
	var ret BaseRet
	err = json.Unmarshal([]byte(res), &ret)
	if err != nil {
		return err
	}
	if ret.Ret == 0 {
		return errors.New(ret.Msg)
	}
	log.Log.Debug("update traffic debug:", ret.Msg)
	return nil
}

func (c *Client) LogNodeOnlineUser(onlineUserCount int) error {
	res, err := c.httpPostNodeOnlineCount(onlineUserCount)
	if err != nil {
		return nil
	}
	var ret BaseRet
	err = json.Unmarshal([]byte(res), &ret)
	if err != nil {
		return err
	}
	if ret.Ret == 0 {
		return UpdateOnlineCountFail
	}
	return nil
}

func (c *Client) UpdateNodeInfo() error {
	uptime, err := system.GetUptime()
	if err != nil {
		log.Log.Error(err)
		uptime = "0"
	}

	load, err := system.GetLoad()
	if err != nil {
		load = "0 0 0"
	}

	res, err := c.httpPostNodeInfo(load, uptime)
	if err != nil {
		return nil
	}
	var ret BaseRet
	err = json.Unmarshal([]byte(res), &ret)
	if err != nil {
		return err
	}
	if ret.Ret == 0 {
		return UpdateOnlineCountFail
	}
	return nil
}
