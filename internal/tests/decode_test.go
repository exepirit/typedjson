package tests

import (
	"github.com/exepirit/typedjson/pkg/typedjson"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnmarshal_StructNotImplementsUnmarshaler(t *testing.T) {
	jsonString := "{\"$type\":\"tests.example\",\"key_1\":\"value\"}"
	dstStruct := example{}

	require.NoError(t, typedjson.Unmarshal([]byte(jsonString), &dstStruct))

	require.Equal(t, "value", dstStruct.Key1)
}
