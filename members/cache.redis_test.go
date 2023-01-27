package members

import (
	"github.com/go-redis/redis/v9"
	"msa-auth/cache"
	"reflect"
	"testing"
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
