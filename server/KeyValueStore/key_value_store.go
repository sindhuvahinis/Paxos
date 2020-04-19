package KeyValueStore

import "sync"

type KeyValue struct {
	keyValueStore map[string]string
	rwm           sync.RWMutex
}

func (k *KeyValue) Initialize() {
	k.keyValueStore = make(map[string]string)
}

func (k *KeyValue) PUT(key, value string) {
	k.rwm.Lock()
	defer k.rwm.Unlock()
	k.keyValueStore[key] = value
}

func (k *KeyValue) GET(key string) (string, bool) {
	value, isKeyExist := k.keyValueStore[key]
	return value, isKeyExist
}

func (k *KeyValue) DELETE(key string) string {
	value, _ := k.GET(key)

	k.rwm.Lock()
	defer k.rwm.Unlock()
	delete(k.keyValueStore, key)
	return value
}

type KeyValueStore interface {
	PUT(key, value string) string
	GET(key string) string
	DELETE(key string) string
}
