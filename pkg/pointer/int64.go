package pointer

func ToInt64Pointer(v int64) *int64 {
	return &v
}

func ToInt64PointerOrNilIfZero(v int64) *int64 {
	if v == 0 {
		return nil
	}
	return &v
}

func FromInt64Pointer(p *int64) int64 {
	return FromInt64PointerOrDefaultIfNil(p, 0)
}

func FromInt64PointerOrDefaultIfNil(v *int64, defaultValue int64) int64 {
	if v == nil {
		return defaultValue
	}
	return *v
}
