package runtime

import (
	//. "utils"
	"reflect"
)

type AppOption func(*options) error
type KeyVal map[string]string
type OsArgs []string

// OsArguments is the list of arguments passed to the CLI at runtime
func OsArguments(args []string) AppOption {
	return func(o *options) error {
		o.osArgs = args
		return nil
	}
}

// Argument is an argument passed programmatically
func Argument(kv KeyVal) AppOption {
	return func(o *options) error {
		var key string
		keys := reflect.ValueOf(kv).MapKeys()
		for i := 0; i < len(keys); i++ {
			key = keys[i].String()
			o.arguments[key] = kv[key]
		}
		return nil
	}
}
