package pointer

func ToInt32Pointer(v int32) *int32 {
	return &v
}

func ToInt32PointerOrNilIfZero(v int32) *int32 {
	if v == 0 {
		return nil
	}
	return &v
}

func FromInt32Pointer(p *int32) int32 {
	return FromInt32PointerOrDefaultIfNil(p, 0)
}

func FromInt32PointerOrDefaultIfNil(v *int32, defaultValue int32) int32 {
	if v == nil {
		return defaultValue
	}
	return *v
}
