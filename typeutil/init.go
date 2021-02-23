package typeutil

type M map[string]interface{}

// StrOf string slice
func StrOf(s ...string) []string {
	return s
}

func ObjOf(i ...interface{}) []interface{} {
	return i
}

//https://github.com/emirpasic/gods
