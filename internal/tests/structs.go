package tests

type example struct {
	Key1 string `json:"key_1"`
}

func (e *example) Type() string {
	return "tests.example"
}

type withIgnoredField struct {
	Ignored bool   `json:"-"`
	Key     string `json:"key"`
}

func (e *withIgnoredField) Type() string {
	return "tests.withIgnoredField"
}
