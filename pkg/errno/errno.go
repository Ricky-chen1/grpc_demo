package errno

type Errno struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

func NewErrno(code int64) Errno {
	errno := Errno{
		Code: code,
		Msg:  CodeTag[int(code)],
	}
	return errno
}

func (e *Errno) Error() string {
	return e.Msg
}
