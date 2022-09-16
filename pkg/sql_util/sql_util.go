package sql_util

func Quote(s string) string {
	return "`" + s + "`"
}
