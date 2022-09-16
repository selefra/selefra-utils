package runtime_util

import (
	"runtime"
)

func Stack() string {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	return string(buf)
}
