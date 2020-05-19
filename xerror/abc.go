package xerror

type XErr interface {
	error
	Err(format string, args ...interface{}) error
	Msg(format string, args ...interface{}) string
	New(format string) XErr
}

type XError interface {
	Recover()
	Panic(err error)
	PanicF(err error, msg string, args ...interface{})
	Wrap(err error) error
	WrapF(err error) error
	PanicErr(d1 interface{}, err error) interface{}
}
