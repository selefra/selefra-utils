package pointer

func ToUInt8Pointer(v uint8) *uint8 {
	return &v
}

func ToUInt8PointerOrNilIfZero(v uint8) *uint8 {
	if v == 0 {
		return nil
	}
	return &v
}

func FromUInt8Pointer(p *uint8) uint8 {
	return FromUInt8PointerOrDefaultIfNil(p, 0)
}

func FromUInt8PointerOrDefaultIfNil(v *uint8, defaultValue uint8) uint8 {
	if v == nil {
		return defaultValue
	}
	return *v
}
