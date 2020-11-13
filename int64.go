package null

// Int64 holds data of nullable int64 value.
type Int64 struct {
	valid bool
	value int64
}

// NewNullInt64 creates new nil int64 value.
func NewNullInt64() Int64 {
	return Int64{
		valid: false,
		value: 0,
	}
}

// NewInt64 creates new int64 value.
func NewInt64(value int64) Int64 {
	return Int64{
		valid: true,
		value: value,
	}
}
