package funcs

import (
	"strings"
	"sync"

	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.CallReg.AddFunction("map", new(Map).init)
}

type Map struct {
	value map[string]interface{}
	mutex sync.RWMutex
}

func (m *Map) init() *Map {
	if m.value == nil {
		m.value = make(map[string]interface{})
	}

	return m
}

func (m *Map) Set(key string, value interface{}) map[string]interface{} {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// separate key with / to create nested map
	vKey := strings.Split(key, "/")

	hWalk := m.value
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

	return m.value
}

func (m *Map) Get(key string, data map[string]interface{}) interface{} {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if data == nil {
		data = m.value
	}

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
