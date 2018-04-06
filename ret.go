package musdk

const (
	TypeSs         = 0 // Default, Shadowsocks
	TypeHttp       = 1 // Http proxy
	TypeForward    = 2 // forward
	TypeVPN        = 3 // vpn
	TypeAnyConnect = 4 // AnyConnect
	TypeV2ray      = 5 // V2ray
)

type BaseRet struct {
	Msg string `json:"msg"`
}

type UserDataRet struct {
	BaseRet
	Data []User `json:"data"`
}

type UserTrafficLog struct {
	UserId int64 `json:"user_id"`
	U      int64 `json:"u"`
	D      int64 `json:"d"`
}
