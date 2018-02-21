package musdk

const (
	TypeSs   = 0 // Default, Shadowsocks
	TypeHttp = 1 // Http proxy
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
