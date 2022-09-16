package json_util

import "encoding/json"

func ToJsonString(v interface{}) string {
	if v == nil {
		return ""
	}
	marshal, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	if len(marshal) == 0 {
		return ""
	}
	return string(marshal)
}

func ToJsonBytes(v interface{}) []byte {
	if v == nil {
		return nil
	}
	marshal, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return marshal
}
