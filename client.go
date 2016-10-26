package musdk

import (
	"errors"
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
	token string
}


