package typedjson

// TypeGetter allows get JSON object type.
// Type is string that stores in `$type` attribute.
type TypeGetter interface {
	Type() string
}
