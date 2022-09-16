package pointer

func ToInt16Pointer(v int16) *int16 {
	return &v
}

func ToInt16PointerOrNilIfZero(v int16) *int16 {
	if v == 0 {
		return nil
	}
	return &v
}

func FromInt16Pointer(p *int16) int16 {
	return FromInt16PointerOrDefaultIfNil(p, 0)
}

func FromInt16PointerOrDefaultIfNil(v *int16, defaultValue int16) int16 {
	if v == nil {
		return defaultValue
	}
	return *v
}
