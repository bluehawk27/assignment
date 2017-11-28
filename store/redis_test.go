package store

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-redis/redis"
)

// TODO make these integration-Tests
var rClient = NewClient()

func TestRedisClient_Ping(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		Client *redis.Client
		Expiry int64
	}

	f := fields{
		Client: rClient.Client,
		Expiry: rClient.Expiry,
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"Redis-Ping", f, "PONG", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RedisClient{
				Client: tt.fields.Client,
				Expiry: tt.fields.Expiry,
			}
			got, err := s.Ping(ctx)
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
	type fields struct {
		Client *redis.Client
		Expiry int64
	}
	f := fields{
		Client: rClient.Client,
		Expiry: rClient.Expiry,
	}
	type args struct {
		key   string
		value interface{}
	}
	a := args{
		key:   "test",
		value: "test value",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Redis-Set", f, a, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RedisClient{
				Client: tt.fields.Client,
				Expiry: tt.fields.Expiry,
			}
			if err := s.Set(ctx, tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("RedisClient.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisClient_Get(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		Client *redis.Client
		Expiry int64
	}
	f := fields{
		Client: rClient.Client,
		Expiry: rClient.Expiry,
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
		fields  fields
		args    args
		want    *Message
		wantErr bool
	}{
		{"Redis-Get", f, a, m, false},
		{"Redis-No-Key", f, a1, m1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RedisClient{
				Client: tt.fields.Client,
				Expiry: tt.fields.Expiry,
			}
			got, err := s.Get(ctx, tt.args.key)
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
