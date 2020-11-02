package typedjson

// TypeGetter allows get JSON object type.
// Type is string that stores in `$type` attribute.
type TypeGetter interface {
	Type() string
}

// Option is object mutator.
// It allow to modify before finally serialization.
type Option interface {
	Mutate(object map[string]interface{})
}
