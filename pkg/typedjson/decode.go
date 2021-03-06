package typedjson

import (
	"encoding/json"
	"reflect"
)

type typeTag struct {
	Tag string `json:"$type"`
}

// Unmarshal deserializes typed JSON object to struct.
// Object type pulls from `$type` JSON attribute and will be compared with value in TypedJson method `Type()`.
func Unmarshal(data []byte, dst TypeGetter) error {
	dType := reflect.ValueOf(dst)
	if dType.Kind() != reflect.Ptr || dType.IsNil() {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(dst)}
	}

	jsonType := typeTag{}
	if err := json.Unmarshal(data, &jsonType); err != nil {
		return &TypeNotDefinedError{}
	}

	expType := dst.Type()
	if jsonType.Tag != expType {
		return &InvalidJsonTypeError{
			Expected: expType,
			Actual:   jsonType.Tag,
		}
	}

	return json.Unmarshal(data, dst)
}
