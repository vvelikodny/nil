package null

// Int32 holds data of nullable int32 value.
type Int32 struct {
	valid bool
	value int32
}

// NewNullInt32 creates new nil int32 value.
func NewNullInt32() Int32 {
	return Int32{
		valid: false,
		value: 0,
	}
}

// NewInt32 creates new int32 value.
func NewInt32(value int32) Int32 {
	return Int32{
		valid: true,
		value: value,
	}
}
