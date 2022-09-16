package reflect_util

import (
	"fmt"
	"reflect"
	"unsafe"
)

func GetStructPtrUnExportedField(source interface{}, fieldName string) reflect.Value {
	v := reflect.ValueOf(source).Elem().FieldByName(fieldName)
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

func SetStructPtrUnExportedStrField(source interface{}, fieldName string, fieldVal interface{}) (err error) {
	v := GetStructPtrUnExportedField(source, fieldName)
	rv := reflect.ValueOf(fieldVal)
	if v.Kind() != rv.Kind() {
		return fmt.Errorf("invalid kind: expected kind %v, got kind: %v", v.Kind(), rv.Kind())
	}
	v.Set(rv)
	return nil
}

func IsNil(x any) bool {

	if x == nil {
		return true
	}

	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}

	return false
}
