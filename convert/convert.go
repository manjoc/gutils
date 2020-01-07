package convert

import "encoding/json"

// StringToBytes convert string to bytes
func StringToBytes(s string) []byte {
	return []byte(s)
}

// BytesToString convert bytes to string
func BytesToString(bytes []byte) string {
	return string(bytes)
}

// JSONToMap json string to map
func JSONToMap(jsonStr string) (map[string]string, error) {
	var dat = map[string]string{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err != nil {
		return dat, err
	}
	return dat, nil
}
