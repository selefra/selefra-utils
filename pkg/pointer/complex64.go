package pointer

func ToComplex64Pointer(v complex64) *complex64 {
	return &v
}

func ToComplex64PointerOrNilIfZero(v complex64) *complex64 {
	if v == 0 {
		return nil
	}
	return &v
}

func FromComplex64Pointer(p *complex64) complex64 {
	return FromComplex64PointerOrDefaultIfNil(p, 0)
}

func FromComplex64PointerOrDefaultIfNil(v *complex64, defaultValue complex64) complex64 {
	if v == nil {
		return defaultValue
	}
	return *v
}
