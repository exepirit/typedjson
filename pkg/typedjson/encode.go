package typedjson

import (
	"encoding/json"
	"reflect"
)

func fieldKey(field reflect.StructField) string {
	key, ok := field.Tag.Lookup("json")
	if !ok {
		key = field.Name
	}
	return key
}

func convertStruct(v interface{}) map[string]interface{} {
	vValue := reflect.ValueOf(v)
	if vValue.Kind() == reflect.Ptr {
		vValue = vValue.Elem()
	}
	vType := vValue.Type()
	m := map[string]interface{}{}
	for i := 0; i < vType.NumField(); i++ {
		field := vType.Field(i)
		key := fieldKey(field)
		m[key] = vValue.Field(i).Interface()
	}
	return m
}

// Marshall convert struct implementing TypeGetter to raw byte array.
// JSON object type will be stored in `$type` tag.
// Additional object modifications makes by Option.
func Marshall(v TypeGetter, options ...Option) ([]byte, error) {
	m := convertStruct(v)
	m["$type"] = v.Type()
	for _, option := range options {
		option.Mutate(m)
	}
	return json.Marshal(m)
}
