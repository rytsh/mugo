package render

import "strings"

var hMap = make(map[string]interface{})

func Hold(key string, value interface{}) map[string]interface{} {
	// separate key with / to create nested map
	vKey := strings.Split(key, "/")

	hWalk := hMap
	for i, v := range vKey {
		if i == len(vKey)-1 {
			hWalk[v] = value
			break
		}

		if _, ok := hWalk[v]; !ok {
			hWalk[v] = make(map[string]interface{})
		}

		if _, ok := hWalk[v].(map[string]interface{}); !ok {
			hWalk[v] = make(map[string]interface{})
		}

		hWalk = hWalk[v].(map[string]interface{})
	}

	return hMap
}

func GetHold() map[string]interface{} {
	return hMap
}

func GetData(key string, data map[string]interface{}) interface{} {
	keyValues := strings.Split(key, "/")

	for i, v := range keyValues {
		if _, ok := data[v]; !ok {
			return nil
		}

		if i == len(keyValues)-1 {
			return data[v]
		}

		var ok bool
		data, ok = data[v].(map[string]interface{})
		if !ok {
			return nil
		}
	}

	return nil
}
