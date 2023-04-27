package helpers

import "encoding/json"

func ByteToMapStringInterface(data []byte) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func ConvertToOtherStruct(from, to interface{}) (interface{}, error) {
	js, err := json.Marshal(from)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(js, &to)

	return to, nil
}
