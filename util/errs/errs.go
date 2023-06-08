package errs

const (
	UnKnownError = -1
)

type Error struct {
	code int
	msg  string
}

func New(code int, msg string) error {
	return &Error{
		code: code,
		msg:  msg,
	}
}

func (e *Error) Error() string {
	return e.msg
}

func Code(err error) int {
	if err == nil {
		return 0
	}
	if e, ok := err.(*Error); ok {
		return e.code
	}
	return UnKnownError
}

func Msg(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(*Error); ok {
		return e.msg
	}
	return err.Error()
}
