package pointer

func ToStringPointer(v string) *string {
	return &v
}

func ToStringPointerOrNilIfEmpty(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

func FromStringPointer(p *string) string {
	return FromStringPointerOrDefaultIfNil(p, "")
}

func FromStringPointerOrDefaultIfNil(v *string, defaultValue string) string {
	if v == nil {
		return defaultValue
	}
	return *v
}
