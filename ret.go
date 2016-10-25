package musdk

type BaseRet struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

type UserDataRet struct {
	BaseRet
	Data []User `json:"data"`
}
