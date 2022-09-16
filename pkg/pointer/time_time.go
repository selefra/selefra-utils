package pointer

import "time"

func NowTimePointer() *time.Time {
	f := time.Now()
	return &f
}

func ToTimePointer(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

func FromTimePointer(t *time.Time) time.Time {
	return FromTimePointerOrDefault(t, time.Time{})
}

func FromTimePointerOrDefault(t *time.Time, defaultValue time.Time) time.Time {
	if t == nil {
		return defaultValue
	}
	return *t
}
