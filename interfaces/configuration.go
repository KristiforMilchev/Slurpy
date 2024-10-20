package interfaces

type Configuration interface {
	Exists() bool
	Create() bool
	Load() bool
	GetKey(key string) interface{}
}
