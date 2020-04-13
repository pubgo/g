package pipe

// IPipe interface
type IPipe interface {
	SortBy(swap interface{}) *_func

	Pipe(fn interface{}) *_func

	Map(fn interface{}) *_func

	Reduce(fn interface{}) *_func

	Any(fn func(v interface{}) bool) bool
	Every(fn func(v interface{}) bool) bool
	Each(fn interface{})

	FilterNil() *_func
	Filter(fn interface{}) *_func

	MustNotNil()

	ToString() string
	ToData(fn ...interface{}) interface{}
	P(tags ...string)
}
