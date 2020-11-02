# typedjson
Tag-typed JSON serialization

typedjson allows you to serialize or deserialize JSON object with type definition stored in `$type` attribute.

## Installation

```shell script
go get -u github.com/exepirit/typedjson
```

## Getting started

Serializable structs must be implement `Type() string` method.

```go
package main

import (
    "fmt"
    "github.com/exepirit/typedjson/pkg/typedjson"
)

type example struct {
	Key1 string `json:"key_1"`
}

func (example) Type() string {
	return "models.example"
}


func main() {
    s := example{
        Key1: "Hello, World!",
    }

    data, _ := typedjson.Marshall(&s)
    fmt.Println(string(data))
    // {"$type":"models.example","key_1":"Hello, World!"}
    _ = typedjson.Unmarshal(data, &s)
}
```
