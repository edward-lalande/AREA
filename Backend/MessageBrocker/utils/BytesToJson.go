package utils

import "encoding/json"

func BytesToJson(v []byte) (jsonMap map[string]interface{}, err error) {
	err = json.Unmarshal(v, &jsonMap)

	if err != nil {
		return nil, err
	}

	return jsonMap, nil
}
