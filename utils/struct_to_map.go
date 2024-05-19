package utils

import "encoding/json"

func StructToMap(object interface{}) (map[string]interface{}, error) {
	var newMap map[string]interface{}
	data, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &newMap)
	if err != nil {
		return nil, err
	}
	return newMap, err
}
