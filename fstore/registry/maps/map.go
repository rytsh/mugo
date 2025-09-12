package maps

import (
	"strings"
	"sync"

	"github.com/rytsh/mugo/fstore"
)

func init() {
	fstore.AddStruct("map", New())
}

type Map struct {
	value map[string]any
	mutex sync.RWMutex
}

func New() *Map {
	return &Map{
		value: make(map[string]any),
	}
}

func (m *Map) Set(key string, value any) map[string]any {
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
			hWalk[v] = make(map[string]any)
		}

		if _, ok := hWalk[v].(map[string]any); !ok {
			hWalk[v] = make(map[string]any)
		}

		hWalk = hWalk[v].(map[string]any)
	}

	return m.value
}

func (m *Map) Get(key string, data map[string]any) any {
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
		data, ok = data[v].(map[string]any)
		if !ok {
			return nil
		}
	}

	return nil
}
