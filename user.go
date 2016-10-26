package musdk

import (
)

type User struct {
	Id             int    `json:"id"`
	Port           int    `json:"port"`
	Passwd         string `json:"passwd"`
	Method         string `json:"method"`
	Enable         int    `json:"enable"`
	TransferEnable int64  `json:"transfer_enable"`
	U              int64  `json:"u"`
	D              int64  `json:"d"`
}


