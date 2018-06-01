package runtime

import (
	//. "utils"
	"reflect"
)

type AppOption func(*options) error
type KeyVal map[string]string
type OsArgs []string

func Version(s string) AppOption {
	return func(o *options) error {
		o.version = s
		return nil
	}
}

func Revision(s string) AppOption {
	return func(o *options) error {
		o.revision = s
		return nil
	}
}

func OsArguments(args []string) AppOption {
	return func(o *options) error {
		o.osArgs = args
		return nil
	}
}

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
