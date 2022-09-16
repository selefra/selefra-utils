package pointer

func TruePointer() *bool {
	b := true
	return &b
}

func FalsePointer() *bool {
	b := false
	return &b
}

func ToBoolPointer(boolValue bool) *bool {
	return &boolValue
}

func ToBoolPointerOrNilIfFalse(boolValue bool) *bool {
	if boolValue {
		return &boolValue
	}
	return nil
}

func FromBoolPointer(boolPointer *bool) bool {
	return FromBoolPointerOrDefault(boolPointer, false)
}

func FromBoolPointerOrDefault(boolPointer *bool, defaultValue bool) bool {
	if boolPointer == nil {
		return defaultValue
	} else {
		return *boolPointer
	}
}
