package options

import (
	"github.com/exepirit/typedjson/pkg/typedjson/options"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMutate_Identify(t *testing.T) {
	mutator := options.Identify{}
	object := map[string]interface{}{
		"key": "value",
	}

	mutator.Mutate(object)

	require.Contains(t, object, "$id")
	require.Equal(t, 0, object["$id"])
}
