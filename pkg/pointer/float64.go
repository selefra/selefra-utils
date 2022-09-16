package pointer

func ToFloat64Pointer(v float64) *float64 {
	return &v
}

func ToFloat64PointerOrNilIfZero(v float64) *float64 {
	if v < 0.0000001 {
		return nil
	}
	return &v
}

func FromFloat64Pointer(p *float64) float64 {
	return FromFloat64PointerOrDefaultIfNil(p, 0)
}

func FromFloat64PointerOrDefaultIfNil(v *float64, defaultValue float64) float64 {
	if v == nil {
		return defaultValue
	}
	return *v
}
