package config

type DB struct {
	ConnectionString string
	Driver           string
	MaxIdle          int
	MaxOpen          int
}
