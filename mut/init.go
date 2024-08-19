package mut

import "sync"

type Map[K comparable, V any] struct {
	Map map[K]V
	Mut *sync.RWMutex `json:"-"`
}

func (m *Map[K, V]) Set(key K, value V) {
	m.Mut.Lock()
	m.Map[key] = value
	m.Mut.Unlock()
}

func (m *Map[K, V]) Get(key K) V {
	m.Mut.RLock()
	defer m.Mut.RUnlock()
	return m.Map[key]
}

func (m *Map[K, V]) GetValueAndState(key K) (data V, ok bool) {
	m.Mut.RLock()
	defer m.Mut.RUnlock()
	data, ok = m.Map[key]
	return
}

func (m *Map[K, V]) Remove(key K) {
	m.Mut.Lock()
	delete(m.Map, key)
	m.Mut.Unlock()
}
