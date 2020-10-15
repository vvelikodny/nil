# null
Golang handy immutable nullable structures to avoid pointers for nil values

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
