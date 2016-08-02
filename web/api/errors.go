package api

type ErrResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (e *ErrResponse) String() string {
	return fmt.Sprintf("<ErrResponse code: %s message: %s>", e.ErrCode, e.ErrMsg)
}

func (e *ErrResponse) Message() string {
	return e.ErrMsg
}

func (e *ErrResponse) Code() int {
	return e.ErrCode
}
