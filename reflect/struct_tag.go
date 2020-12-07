package reflect

import (
	"net/url"
	"reflect"
	"strconv"
)

// ToQuery converts a struct to query strings.  A simple example for parsing
// struct tags.
func ToQuery(data interface{}) url.Values {
	if data == nil {
		return nil
	}

	ret := make(url.Values)

	elem := reflect.ValueOf(data).Elem()
	for i := 0; i < elem.NumField(); i++ {
		vf := elem.Field(i)
		tf := elem.Type().Field(i)

		key := tf.Tag.Get("form")
		if key == "" || key == "-" {
			continue
		}

		var v string

		switch vf.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v = strconv.FormatInt(vf.Int(), 10)
		case reflect.String:
			v = vf.String()
		default:
			continue
		}

		if v == "" {
			continue
		}

		ret[key] = append(ret[key], v)
	}

	return ret
}
