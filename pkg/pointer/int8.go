package pointer

func ToInt8Pointer(v int8) *int8 {
	return &v
}

func ToInt8PointerOrNilIfZero(v int8) *int8 {
	if v == 0 {
		return nil
	}
	return &v
}

func FromInt8Pointer(p *int8) int8 {
	return FromInt8PointerOrDefaultIfNil(p, 0)
}

func FromInt8PointerOrDefaultIfNil(v *int8, defaultValue int8) int8 {
	if v == nil {
		return defaultValue
	}
	return *v
}
