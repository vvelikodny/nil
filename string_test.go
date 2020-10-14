package null_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vvelikodny/null"
)

func TestNilString(t *testing.T) {
	str1 := null.NewNullString()
	require.True(t, str1.Null())
	require.Equal(t, "nil", fmt.Sprint(str1))

	str2 := null.NewString("value")
	require.False(t, str2.Null())
	require.Equal(t, "value", fmt.Sprint(str2))
}

type Json struct {
	Field null.String `json:"field"`
}

func TestJsonWithNullString(t *testing.T) {
	var b bytes.Buffer
	require.NoError(t, json.NewEncoder(&b).Encode(Json{
		Field: null.NewNullString(),
	}))
	require.Equal(t, `{"field":null}
`, b.String())
}
