package tests

import (
	"encoding/json"
	"github.com/exepirit/typedjson/pkg/typedjson"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMarshall_StructPointer(t *testing.T) {
	s := example{Key1: "value"}

	marshalled, err := typedjson.Marshall(&s)

	require.NoError(t, err)
	m := map[string]string{}
	require.NoError(t, json.Unmarshal(marshalled, &m))

	require.JSONEq(t, `{"$type":"tests.example","key_1":"value"}`, string(marshalled))
}
