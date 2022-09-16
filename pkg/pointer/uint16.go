package pointer

func ToUInt16Pointer(v uint16) *uint16 {
	return &v
}

func ToUInt16PointerOrNilIfZero(v uint16) *uint16 {
	if v == 0 {
		return nil
	}
	return &v
}

func FromUInt16Pointer(p *uint16) uint16 {
	return FromUInt16PointerOrDefaultIfNil(p, 0)
}

func FromUInt16PointerOrDefaultIfNil(v *uint16, defaultValue uint16) uint16 {
	if v == nil {
		return defaultValue
	}
	return *v
}
