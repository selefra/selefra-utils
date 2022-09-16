package pointer

func ToComplex128Pointer(v complex128) *complex128 {
	return &v
}

func ToComplex128PointerOrNilIfZero(v complex128) *complex128 {
	if v == 0 {
		return nil
	}
	return &v
}

func FromComplex128Pointer(p *complex128) complex128 {
	return FromComplex128PointerOrDefaultIfNil(p, 0)
}

func FromComplex128PointerOrDefaultIfNil(v *complex128, defaultValue complex128) complex128 {
	if v == nil {
		return defaultValue
	}
	return *v
}
