[![Build Status](https://travis-ci.com/vvelikodny/nil.svg?branch=main)](https://travis-ci.com/vvelikodny/nil)

# nil
Golang immutable nilable structures to avoid pointers for nil values.

# status

development in progress

# Usage

```golang
package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/vvelikodny/nil"
)

type Json struct {
	Field nil.String `json:"field"`
}

func main() {
	str1 := nil.NewNilString()
	str2 := nil.NewString("value")

	var b bytes.Buffer
	json.NewEncoder(&b).Encode(Json{
		Field: null.NewNilString(),
	})
	// `{"field":null}`
	fmt.Println(b.String())

    var b *bool
    value := nil.FromBoolPtr(b)
    // `nil? true`
    fmt.Println("nil? ", value.Nil())

    b = true
	value := nil.FromBoolPtr(&b)
    // `nil? false`
	fmt.Println("nil?", value.Nil())
}
```