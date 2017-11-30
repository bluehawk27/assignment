package store

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

type FakeRedis struct {
	PingFunc func(context.Context) (string, error)
	SetFunc  func(context.Context, string, interface{}) error
	GetFunc  func(context.Context, string) (*Message, error)
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
func (r FakeRedis) Get(ctx context.Context, key string) (*Message, error) {
	var m *Message
	// if r.GetFunc != nil {
	// 	return r.GetFunc(ctx, key)
	// }
	if key == "test no Key" {
		m = &Message{
			Key:   key,
			Value: "Key Does Not Exist",
		}
		return m, nil
	}
	if key == "test" {
		m = &Message{
			Key:   key,
			Value: "test value",
		}
		return m, nil
	}

	return nil, fmt.Errorf("Get %s Error", key)
}

func TestRedisClient_Ping(t *testing.T) {
	ctx := context.Background()
	f := FakeRedis{}
	tests := []struct {
		name    string
		fields  FakeRedis
		want    string
		wantErr bool
	}{
		{"Redis-Ping", f, "PONG", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.Ping(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisClient.Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RedisClient.Ping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisClient_Set(t *testing.T) {
	ctx := context.Background()
	f := FakeRedis{}
	type args struct {
		Key   string
		Value interface{}
	}
	a := args{
		Key:   "test",
		Value: "test value",
	}
	tests := []struct {
		name    string
		fields  FakeRedis
		args    args
		wantErr bool
	}{
		{"Redis-Set", f, a, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := f.Set(ctx, tt.args.Key, tt.args.Value)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisClient.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisClient_Get(t *testing.T) {
	ctx := context.Background()
	f := FakeRedis{}
	type args struct {
		Key string
	}
	a := args{
		Key: "test",
	}
	a1 := args{
		Key: "test no Key",
	}
	m := &Message{
		Key:   "test",
		Value: "test value",
	}
	m1 := &Message{
		Key:   "test no Key",
		Value: "Key Does Not Exist",
	}
	tests := []struct {
		name    string
		fields  FakeRedis
		args    args
		want    *Message
		wantErr bool
	}{
		{"Redis-Get", f, a, m, false},
		{"Redis-No-Key", f, a1, m1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.Get(ctx, tt.args.Key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RedisClient.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO make these integration-Tests
// var rClient = NewClient()

// func TestRedisClient_Ping(t *testing.T) {
// 	ctx := context.Background()
// 	type fields struct {
// 		Client *redis.Client
// 		Expiry int64
// 	}

// 	f := fields{
// 		Client: rClient.Client,
// 		Expiry: rClient.Expiry,
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    string
// 		wantErr bool
// 	}{
// 		{"Redis-Ping", f, "PONG", false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &RedisClient{
// 				Client: tt.fields.Client,
// 				Expiry: tt.fields.Expiry,
// 			}
// 			got, err := s.Ping(ctx)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RedisClient.Ping() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("RedisClient.Ping() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestRedisClient_Set(t *testing.T) {
// 	ctx := context.Background()
// 	type fields struct {
// 		Client *redis.Client
// 		Expiry int64
// 	}
// 	f := fields{
// 		Client: rClient.Client,
// 		Expiry: rClient.Expiry,
// 	}
// 	type args struct {
// 		key   string
// 		value interface{}
// 	}
// 	a := args{
// 		key:   "test",
// 		value: "test value",
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		{"Redis-Set", f, a, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &RedisClient{
// 				Client: tt.fields.Client,
// 				Expiry: tt.fields.Expiry,
// 			}
// 			if err := s.Set(ctx, tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
// 				t.Errorf("RedisClient.Set() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestRedisClient_Get(t *testing.T) {
// 	ctx := context.Background()
// 	type fields struct {
// 		Client *redis.Client
// 		Expiry int64
// 	}
// 	f := fields{
// 		Client: rClient.Client,
// 		Expiry: rClient.Expiry,
// 	}
// 	type args struct {
// 		key string
// 	}
// 	a := args{
// 		key: "test",
// 	}
// 	a1 := args{
// 		key: "test no Key",
// 	}
// 	m := &Message{
// 		Key:   "test",
// 		Value: "test value",
// 	}
// 	m1 := &Message{
// 		Key:   "test no Key",
// 		Value: "Key Does Not Exist",
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *Message
// 		wantErr bool
// 	}{
// 		{"Redis-Get", f, a, m, false},
// 		{"Redis-No-Key", f, a1, m1, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &RedisClient{
// 				Client: tt.fields.Client,
// 				Expiry: tt.fields.Expiry,
// 			}
// 			got, err := s.Get(ctx, tt.args.key)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RedisClient.Get() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("RedisClient.Get() = %v, want %v", got, tt.want)
// 			}
// 		})
// }
