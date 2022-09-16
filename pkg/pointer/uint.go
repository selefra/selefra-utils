package pointer

func ToUIntPointer(v uint) *uint {
	return &v
}

func ToUIntPointerOrNilIfZero(v uint) *uint {
	if v == 0 {
		return nil
	}
	return &v
}

func FromUIntPointer(p *uint) uint {
	return FromUIntPointerOrDefaultIfNil(p, 0)
}

func FromUIntPointerOrDefaultIfNil(v *uint, defaultValue uint) uint {
	if v == nil {
		return defaultValue
	}
	return *v
}
