package main

import "fmt"

type mapStore struct {
	store map[string]string
}

func newMapStore() *mapStore {
	return &mapStore{
		store: make(map[string]string),
	}
}

func (m *mapStore) setMap(key, value string) {
	m.store[key] = value
}

func (m *mapStore) getMap(key string) (string, bool) {
	value, exists := m.store[key]
	return value, exists

}

func (m *mapStore) deleteMap(key string) {
	delete(m.store, key)
}

func (m *mapStore) ListAllKeys() []string {
	list := make([]string, 0, len(m.store))
	for keys := range m.store {
		list = append(list, keys)
	}
	return list
}

func main() {
	finalMap := newMapStore()
	finalMap.setMap("name", "Rahul")
	finalMap.setMap("age", "22")
	finalMap.setMap("city", "TVM")
	finalMap.getMap("name")
	finalMap.getMap("address")
	// keys := finalMap.ListAllKeys()
	// fmt.Println("The keys in the store:", keys)
	finalMap.deleteMap("age")
	keys := finalMap.ListAllKeys()
	fmt.Println("The keys in the store:", keys)
}
