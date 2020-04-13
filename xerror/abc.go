package xerror

// IErr err interface
type IErr interface {
	error

	P()
	M(k string, v interface{})
	Caller(caller ...string)
	Err() error
	SetErr(err interface{}, args ...interface{})
	Is(err error) bool
	As(err error) bool
	StackTrace() *_Err1

	next() IErr
	node(IErr)
	_err() *_Err1
	_err1() []*_Err1
}
