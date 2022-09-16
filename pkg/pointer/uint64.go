package pointer

func ToUInt64Pointer(v uint64) *uint64 {
	return &v
}

func ToUInt64PointerOrNilIfZero(v uint64) *uint64 {
	if v == 0 {
		return nil
	}
	return &v
}

func FromUInt64Pointer(p *uint64) uint64 {
	return FromUInt64PointerOrDefaultIfNil(p, 0)
}

func FromUInt64PointerOrDefaultIfNil(v *uint64, defaultValue uint64) uint64 {
	if v == nil {
		return defaultValue
	}
	return *v
}
