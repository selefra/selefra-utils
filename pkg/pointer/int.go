package pointer

func ToIntPointer(v int) *int {
	return &v
}

func ToIntPointerOrNilIfZero(v int) *int {
	if v == 0 {
		return nil
	}
	return &v
}

func FromIntPointer(p *int) int {
	return FromIntPointerOrDefaultIfNil(p, 0)
}

func FromIntPointerOrDefaultIfNil(v *int, defaultValue int) int {
	if v == nil {
		return defaultValue
	}
	return *v
}
