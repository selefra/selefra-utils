package id_util

import "strings"
import uuid2 "github.com/satori/go.uuid"

// RandomUUID A UUID V4 string is randomly generated
func RandomUUID() string {
	return uuid2.NewV4().String()
}

// RandomId Generates a globally unique 32-bit character ID, which is essentially a UUID without the - sign
func RandomId() string {
	return strings.ReplaceAll(uuid2.NewV4().String(), "-", "")
}
