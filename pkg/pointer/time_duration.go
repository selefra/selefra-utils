package pointer

import "time"

func ToDurationPointer(v time.Duration) *time.Duration {
	return &v
}

func ToDurationPointerOrNilIfZero(v time.Duration) *time.Duration {
	if v == 0 {
		return nil
	}
	return &v
}

func FromDurationPointer(p *time.Duration) time.Duration {
	return FromDurationPointerOrDefaultIfNil(p, 0)
}

func FromDurationPointerOrDefaultIfNil(v *time.Duration, defaultValue time.Duration) time.Duration {
	if v == nil {
		return defaultValue
	}
	return *v
}
