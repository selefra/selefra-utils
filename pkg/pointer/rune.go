package pointer

func ToRunePointer(v rune) *rune {
	return &v
}

func ToRunePointerOrNilIfZero(v rune) *rune {
	if v == 0 {
		return nil
	}
	return &v
}

func FromRunePointer(p *rune) rune {
	return FromRunePointerOrDefaultIfNil(p, 0)
}

func FromRunePointerOrDefaultIfNil(v *rune, defaultValue rune) rune {
	if v == nil {
		return defaultValue
	}
	return *v
}
