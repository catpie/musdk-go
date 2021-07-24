package musdk

import "fmt"

const (
	TypeSs         = 0 // Default, Shadowsocks
	TypeHttp       = 1 // Http proxy
	TypeForward    = 2 // forward
	TypeVPN        = 3 // vpn
	TypeAnyConnect = 4 // AnyConnect
	TypeV2ray      = 5 // V2ray
)

type Node struct {
	ID                    int32  `json:"id"`
	Name                  string `json:"name"`
	ServerMonitorAddr     string `json:"server_monitor_addr"`
	ServerMonitorPort     int32  `json:"server_monitor_port"`
	ServerMonitorGrpcPort int32  `json:"server_monitor_grpc_port"`
	Server                string `json:"server"`
}

func (n *Node) GetMonitorWsAddr() string {
	var port int32 = 8080
	if n.ServerMonitorPort != 0 {
		port = n.ServerMonitorPort
	}
	return fmt.Sprintf("ws://%s:%d/ws", n.Server, port)
}

func (n *Node) GetMonitorGrpcAddr() string {
	var port int32 = 8090
	if n.ServerMonitorGrpcPort != 0 {
		port = n.ServerMonitorGrpcPort
	}
	return fmt.Sprintf("%s:%d", n.Server, port)
}

type BaseRet struct {
	Msg string `json:"msg"`
}

type UserDataRet struct {
	BaseRet
	Data []User `json:"data"`
}

type NodeDataRet struct {
	BaseRet
	Data []Node `json:"data"`
}

type UserTrafficLog struct {
	UserId int64 `json:"user_id"`
	U      int64 `json:"u"`
	D      int64 `json:"d"`
}

type OnlineInfo struct {
	Count int64 `json:"count"`
}
