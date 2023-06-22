package util
import "fmt"

type ErrNo struct {
	ErrMsg string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_msg=%s", e.ErrMsg)
}

func NewErrNo(msg string) ErrNo {
	return ErrNo{
		ErrMsg: msg,
	}
}