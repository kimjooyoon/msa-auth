package api

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestOk(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 gin.H
	}{
		{"success", 200, gin.H{"message": "success"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Ok()
			if got != tt.want {
				t.Errorf("Ok() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Ok() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOkWithMessage(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 gin.H
	}{
		{"success", args{"hello world"}, 200, gin.H{"message": "hello world"}},
		{"success", args{"안녕하세요."}, 200, gin.H{"message": "안녕하세요."}},
		{"success", args{"test"}, 200, gin.H{"message": "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OkWithMessage(tt.args.message)
			if got != tt.want {
				t.Errorf("OkWithMessage() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OkWithMessage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestServerError(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 gin.H
	}{
		{"success", 500, gin.H{"message": "fail"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ServerError()
			if got != tt.want {
				t.Errorf("ServerError() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ServerError() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOkWithToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 gin.H
	}{
		{"success", args{"hello world"}, 200, gin.H{"token": "hello world"}},
		{"success", args{"안녕하세요."}, 200, gin.H{"token": "안녕하세요."}},
		{"success", args{"test"}, 200, gin.H{"token": "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OkWithToken(tt.args.token)
			if got != tt.want {
				t.Errorf("OkWithToken() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("OkWithToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
