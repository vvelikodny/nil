[![Build Status](https://travis-ci.com/vvelikodny/null.svg?branch=main)](https://travis-ci.com/vvelikodny/null)

# null
Golang nullable immutable structures to avoid pointers for nil values

# status

development in progress

# Usage

```golang
package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/vvelikodny/null"
)

type Json struct {
	Field null.String `json:"field"`
}

func main() {
	str1 := null.NewNullString()
	str2 := null.NewString("value")

	var b bytes.Buffer
	json.NewEncoder(&b).Encode(Json{
		Field: null.NewNullString(),
	})
	// `{"field":null}`
	fmt.Println(b.String())
}
```
