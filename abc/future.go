package abc

import "reflect"

type IPromise interface {
	Await() chan FutureValue
	Map(fn interface{}) Value
}

type Future interface {
	Yield(data interface{})                  // async
	YieldFn(val FutureValue, fn interface{}) // block
}

type FutureValue interface {
	Expect(format string, a ...interface{})
	Err() error
	String() string
	Get() interface{}
	Raw() []reflect.Value
	Value(fn interface{}) error
}

type Value interface {
	Assert(format string, a ...interface{})
	Err() error
	Value() interface{}
}
