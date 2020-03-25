package util

const (
	SUCCESS        = 0
	ERROR          = 1
	INVALID_PARAMS = 2
)

var MsgFlags = map[uint32]string{
	SUCCESS:        "ok",
	ERROR:          "内部错误",
	INVALID_PARAMS: "请求参数错误",
}

func GetMsg(code uint32) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
