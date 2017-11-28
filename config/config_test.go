package config

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Init"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
		})
	}
}

func TestGetCCacheConfig(t *testing.T) {
	c := CCache{
		Capacity: int64(10),
		Expiry:   int64(60),
	}
	tests := []struct {
		name string
		want CCache
	}{
		{"CCache-Config", c},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCCacheConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCCacheConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRedisConfig(t *testing.T) {
	r := Redis{
		Host:     "localhost",
		Port:     "6379",
		Password: "",
		DB:       0,
		Expiry:   int64(60),
	}
	tests := []struct {
		name string
		want Redis
	}{
		{"Redis-Config", r},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedisConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRedisConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProxyConfig(t *testing.T) {
	p := Proxy{
		Host: "127.0.0.1",
		Port: "8082",
	}
	tests := []struct {
		name string
		want Proxy
	}{
		{"Proxy-Config", p},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProxyConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProxyConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetProxyConnectionString(t *testing.T) {
	s := "127.0.0.1:8082"
	tests := []struct {
		name string
		want string
	}{
		{"Proxy-String", s},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetProxyConnectionString(); got != tt.want {
				t.Errorf("GetProxyConnectionString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRedisConnectionString(t *testing.T) {
	s := "localhost:6379"
	tests := []struct {
		name string
		want string
	}{
		{"Redis-String", s},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedisConnectionString(); got != tt.want {
				t.Errorf("GetRedisConnectionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
