package config

type Redis struct {
	Host     string
	Port     string
	Password string
	DB       int
	Expiry   int64
}

type Proxy struct {
	Host string
	Port string
}

type CCache struct {
	Capacity int64
	Expiry   int64
}
