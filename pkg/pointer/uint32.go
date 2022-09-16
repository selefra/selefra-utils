package pointer

func ToUInt32Pointer(v uint32) *uint32 {
	return &v
}

func ToUInt32PointerOrNilIfZero(v uint32) *uint32 {
	if v == 0 {
		return nil
	}
	return &v
}

func FromUInt32Pointer(p *uint32) uint32 {
	return FromUInt32PointerOrDefaultIfNil(p, 0)
}

func FromUInt32PointerOrDefaultIfNil(v *uint32, defaultValue uint32) uint32 {
	if v == nil {
		return defaultValue
	}
	return *v
}
