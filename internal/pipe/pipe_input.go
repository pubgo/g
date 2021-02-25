package pipe

import "reflect"

// ArrayOf input array
func ArrayOf(ps interface{}) IPipe {
	_d := reflect.ValueOf(ps)

	var _ps []reflect.Value
	for i := 0; i < _d.Len(); i++ {
		_ps = append(_ps, _d.Index(i))
	}
	return &_func{
		params: _ps,
	}
}

// SliceOf input slice
func SliceOf(ps ...interface{}) IPipe {
	var vs = make([]reflect.Value, len(ps))
	for i, v := range ps {
		vs[i] = reflect.ValueOf(v)
	}
	return &_func{
		params: vs,
	}
}

// SortBy input array
func SortBy(data interface{}, swap interface{}) IPipe {
	return ArrayOf(data).SortBy(swap)
}

// Map input array
func Map(data interface{}, fn interface{}) IPipe {
	return ArrayOf(data).Map(fn)
}

// Reduce input array
func Reduce(data interface{}, code string) IPipe {
	return ArrayOf(data).Reduce(code)
}

// Any input array
func Any(data interface{}, fn func(v interface{}) bool) bool {
	return ArrayOf(data).Any(fn)
}

// Every input array
func Every(data interface{}, fn func(v interface{}) bool) bool {
	return ArrayOf(data).Every(fn)
}

// Each input array
func Each(data interface{}, fn interface{}) {
	ArrayOf(data).Each(fn)
}

// FilterNil input array
func FilterNil(data interface{}) IPipe {
	return ArrayOf(data).FilterNil()
}

// Filter input array
func Filter(data interface{}, fn interface{}) IPipe {
	return ArrayOf(data).Filter(fn)
}

// Pipe input array
func Pipe(data interface{}, fn interface{}) IPipe {
	return ArrayOf(data).Pipe(fn)
}
