package md5_util

import (
	"crypto/md5"
	"fmt"
)

func Md5String(s string) (string, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(s))
	if err != nil {
		return "", err
	}
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum), nil
}
