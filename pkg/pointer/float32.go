package pointer

func ToFloat32Pointer(v float32) *float32 {
	return &v
}

func ToFloat32PointerOrNilIfZero(v float32) *float32 {
	if v < 0.0000001 {
		return nil
	}
	return &v
}

func FromFloat32Pointer(p *float32) float32 {
	return FromFloat32PointerOrDefaultIfNil(p, 0)
}

func FromFloat32PointerOrDefaultIfNil(v *float32, defaultValue float32) float32 {
	if v == nil {
		return defaultValue
	}
	return *v
}
