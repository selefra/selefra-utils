package string_util

import "strings"

type StringBuilder struct {
	sb strings.Builder
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{sb: strings.Builder{}}
}

func (x *StringBuilder) WriteString(s string) *StringBuilder {
	x.sb.WriteString(s)
	return x
}

func (x *StringBuilder) Write(bytes []byte) *StringBuilder {
	x.sb.Write(bytes)
	return x
}

func (x *StringBuilder) Grow(n int) *StringBuilder {
	x.sb.Grow(n)
	return x
}

func (x *StringBuilder) Reset() *StringBuilder {
	x.sb.Reset()
	return x
}

func (x *StringBuilder) GetInnerBuilder() *strings.Builder {
	return &x.sb
}

func (x *StringBuilder) Len() int {
	return x.sb.Len()
}

func (x *StringBuilder) Cap() int {
	return x.sb.Cap()
}

func (x *StringBuilder) String() string {
	return x.sb.String()
}
