package coproduct

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func DecodeTaggedUnion(d []byte, tf, vf reflect.StructField, i interface{}, m map[string]reflect.Type) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(d, &raw); err != nil {
		return err
	}
	ty, ok := tf.Tag.Lookup("json")
	if !ok {
		ty = tf.Name
	}
	rawTag, ok := raw[ty]
	if !ok {
		return fmt.Errorf("field not found: %s", ty)
	}
	tag, ok := rawTag.(string)
	if !ok {
		return fmt.Errorf("tag not type of string: %v", rawTag)
	}
	t, ok := m[tag]
	if !ok {
		return fmt.Errorf("type not found: %s", tag)
	}
	v := reflect.New(t).Interface()
	val, ok := vf.Tag.Lookup("json")
	if !ok {
		val = vf.Name
	}
	dd, err := json.Marshal(raw[val])
	if err != nil {
		return err
	}
	if err := json.Unmarshal(dd, v); err != nil {
		return err
	}
	ival := reflect.ValueOf(i)
	ival.Elem().FieldByName(tf.Name).Set(reflect.ValueOf(tag))
	ival.Elem().FieldByName(vf.Name).Set(reflect.ValueOf(v))
	return nil
}
