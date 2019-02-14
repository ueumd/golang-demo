package errno

import "fmt"

/**
Errno 用来自定义错误Code，比如包含：Code和Message信息。
Err是一个 error 类型的结构体，代表一个错误。
 */
type Errno struct {
	Code    int
	Message string
}

// 实现了error接口中的Error方法
func (err Errno) Error() string {
	return err.Message
}

// Err represents an error
type Err struct {
	Code    int
	Message string
	Err     error
}

// 新建定制的错误
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

// 解析定制的错误
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	// 类型断言
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		fmt.Println("222222222222")
		fmt.Printf("typed: %T \n", typed)
		fmt.Printf("Errno: %T \n", &Errno{})
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
