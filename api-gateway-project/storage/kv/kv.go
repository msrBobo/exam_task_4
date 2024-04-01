package kv

import "exam_task_4/api-gateway-project/storage"

type KV interface {
	Set(key string, value *storage.User) error
	Get(key string) (*storage.User, error)
	Delete(key string) error
}

var inst KV

func Init(store KV) {
	inst = store
}

func Set(key string, value *storage.User) error {
	return inst.Set(key, value)
}
func Get(key string) (*storage.User, error) {
	return inst.Get(key)
}
func Delete(key string) error {
	return inst.Delete(key)
}
