package coproduct

import (
	"encoding/json"
	"reflect"
)

func DecodeUntaggedUnion(d []byte, i interface{}, types ...reflect.Type) error {
	var err error
	for _, ty := range types {
		v := reflect.New(ty).Interface()
		if err = json.Unmarshal(d, v); err == nil {
			ival := reflect.ValueOf(i)
			ival.Elem().Set(reflect.ValueOf(v))
			break
		}
	}
	return err
}
