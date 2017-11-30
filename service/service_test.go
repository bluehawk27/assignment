package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/bluehawk27/assignment/cache"
	"github.com/bluehawk27/assignment/store"
)

//TODO MOCK the Redis Service
var servType = NewService()

type FakeRedis struct {
	PingFunc func(context.Context) (string, error)
	SetFunc  func(context.Context, string, interface{}) error
	GetFunc  func(context.Context, string) (*store.Message, error)
}

// Ping : FakeRedis Ping
func (r FakeRedis) Ping(ctx context.Context) (string, error) {
	return "PONG", nil
}

// Set : FakeRedis Set
func (r FakeRedis) Set(ctx context.Context, key string, value interface{}) error {
	if r.SetFunc != nil {
		return fmt.Errorf("Error Setting")
	}
	return nil
}

// Get : FakeRedis Get
func (r FakeRedis) Get(ctx context.Context, key string) (*store.Message, error) {
	var m *store.Message
	if r.GetFunc != nil {
		return r.GetFunc(ctx, key)
	}
	if key == "test no Key" {
		m = &store.Message{
			Key:   key,
			Value: "Key Does Not Exist",
		}
		return m, nil
	}

	return nil, fmt.Errorf("Get %s Error", key)
}

func TestNewService(t *testing.T) {
	tests := []struct {
		name string
		want ServiceType
	}{
		{"Service-New", servType},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := servType; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Ping(t *testing.T) {
	ctx := context.Background()
	fr := FakeRedis{}
	type fields struct {
		Store FakeRedis
		Cache cache.Cache
	}
	f := fields{
		Store: fr,
		Cache: servType.Cache,
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"Service-Ping", f, "PONG", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: tt.fields.Store,
				Cache: tt.fields.Cache,
			}
			got, err := s.Ping(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.Ping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Add(t *testing.T) {
	ctx := context.Background()
	fr := FakeRedis{}
	type fields struct {
		Store FakeRedis
		Cache cache.Cache
	}
	f := fields{
		Store: fr,
		Cache: servType.Cache,
	}
	type args struct {
		key  string
		body []byte
	}
	a := args{
		key:  "test",
		body: []byte("test Body"),
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Service-Add", f, a, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: tt.fields.Store,
				Cache: tt.fields.Cache,
			}
			if err := s.Add(ctx, tt.args.key, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("Service.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_Get(t *testing.T) {
	ctx := context.Background()
	fr := FakeRedis{}

	type fields struct {
		Store FakeRedis
		Cache cache.Cache
	}
	f := fields{
		Store: fr,
		Cache: servType.Cache,
	}
	type args struct {
		key string
	}
	a := args{
		key: "test",
	}
	a1 := args{
		key: "test no Key",
	}
	m := &store.Message{
		Key:   a.key,
		Value: "test Body",
	}
	m1 := &store.Message{
		Key:   a1.key,
		Value: "Key Does Not Exist",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *store.Message
		wantErr bool
	}{
		{"Service-Get", f, a, m, false},
		{"Service-NO-KEY", f, a1, m1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: tt.fields.Store,
				Cache: tt.fields.Cache,
			}
			got, err := s.Get(ctx, tt.args.key)
			// fmt.Println(got)
			// fmt.Println(err)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
