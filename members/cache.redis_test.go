package members

import (
	"context"
	"github.com/go-redis/redis/v9"
	"msa-auth/cache"
	"reflect"
	"testing"
	"time"
)

func TestNewRedis(t *testing.T) {
	type args struct {
		r   *redis.Client
		ctx cache.Context
	}
	tests := []struct {
		name string
		args args
		want R
	}{
		{"success", args{}, NewRedis(nil, nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedis(tt.args.r, tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockRdsClient struct{}

func (m mockRdsClient) Set(context.Context, string, interface{}, time.Duration) *redis.StatusCmd {
	return &redis.StatusCmd{}
}
func (m mockRdsClient) Get(context.Context, string) *redis.StringCmd {
	return nil
}

func TestRC_Logout(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		rdb RdsClient
		ctx cache.Context
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"success", fields{mockRdsClient{}, ctx}, args{""},
			false},
		{"success", fields{mockRdsClient{}, ctx}, args{""},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RC{
				rdb: tt.fields.rdb,
				ctx: tt.fields.ctx,
			}
			if err := r.Logout(tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
