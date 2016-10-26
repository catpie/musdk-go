package musdk

type BaseRet struct {
	Msg string `json:"msg"`
}

type UserDataRet struct {
	BaseRet
	Data []User `json:"data"`
}

type UserLog struct {
	UserId int64 `json:"user_id"`
	U      int64 `json:"u"`
	D      int64 `json:"d"`
}
