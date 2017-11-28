package cache

import (
	"reflect"
	"testing"

	"github.com/karlseguin/ccache"
)

var cache = NewCache()

func TestNewCache(t *testing.T) {
	tests := []struct {
		name string
		want *Cache
	}{
		{"success", cache},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cache; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_Set(t *testing.T) {
	type fields struct {
		Cache  *ccache.Cache
		Expiry int64
	}
	f := fields{
		Cache:  cache.Cache,
		Expiry: cache.Expiry,
	}
	type args struct {
		key   string
		value interface{}
	}

	a := args{
		key:   "key",
		value: "value",
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Cache-Add", f, a},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				Cache:  tt.fields.Cache,
				Expiry: tt.fields.Expiry,
			}
			c.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestCache_Get(t *testing.T) {
	type fields struct {
		Cache  *ccache.Cache
		Expiry int64
	}

	f := fields{
		Cache:  cache.Cache,
		Expiry: cache.Expiry,
	}

	type args struct {
		key string
	}

	a1 := args{
		key: "key3",
	}
	var ret1 interface{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"Cache Miss", f, a1, ret1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				Cache:  tt.fields.Cache,
				Expiry: tt.fields.Expiry,
			}
			got, err := c.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
