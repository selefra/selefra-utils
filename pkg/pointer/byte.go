package pointer

func ToBytePointer(v byte) *byte {
	return &v
}

func ToBytePointerOrNilIfZero(v byte) *byte {
	if v == 0 {
		return nil
	}
	return &v
}

func FromBytePointer(p *byte) byte {
	return FromBytePointerOrDefaultIfNil(p, 0)
}

func FromBytePointerOrDefaultIfNil(v *byte, defaultValue byte) byte {
	if v == nil {
		return defaultValue
	}
	return *v
}
