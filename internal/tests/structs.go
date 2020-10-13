package tests

type example struct {
	Key1 string `json:"key_1"`
}

func (e *example) Type() string {
	return "tests.example"
}
