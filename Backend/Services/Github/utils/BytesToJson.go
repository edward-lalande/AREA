package utils

import "encoding/json"

func BytesToJson(v []byte) (jsonMap map[string]interface{}) {
	json.Unmarshal(v, &jsonMap)

	return jsonMap
}
