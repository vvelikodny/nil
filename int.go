package null

// Int holds data of nullable int value.
type Int struct {
	valid bool
	value int
}

// NewNullInt creates new nil int value.
func NewNullInt() Int {
	return Int{
		valid: false,
		value: 0,
	}
}

// NewInt creates new int value.
func NewInt(value int) Int {
	return Int{
		valid: true,
		value: value,
	}
}
