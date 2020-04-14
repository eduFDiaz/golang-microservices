package domain

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang-microservices/mvc/utils"
)

func TestGetUser(t *testing.T) {
	type args struct {
		userID int64
	}
	tests := []struct {
		name  string
		args  args
		want  *User
		want1 *utils.ApplicationError
	}{
		{
			name: "Test Get user id with 1",
			args: args{int64(1)},
			want: &User{
				ID:        1,
				FirstName: "Eduardo",
				LastName:  "Fernandez",
				Email:     "fenandez9000@gmail.com",
			},
			want1: nil,
		},
		{
			name: "Test Get user id with 123",
			args: args{int64(123)},
			want: nil,
			want1: &utils.ApplicationError{
				Message:    fmt.Sprintf("User with id = %v not found", int64(123)),
				StatusCode: http.StatusNotFound,
				Code:       "not found",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetUser(tt.args.userID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func BenchmarkGetUser(b *testing.B) {
	type args struct {
		userID int64
	}
	tests := []struct {
		name  string
		args  args
		want  *User
		want1 *utils.ApplicationError
	}{
		{
			name: "Test Get user id with 1",
			args: args{int64(1)},
			want: &User{
				ID:        1,
				FirstName: "Eduardo",
				LastName:  "Fernandez",
				Email:     "fenandez9000@gmail.com",
			},
			want1: nil,
		},
		{
			name: "Test Get user id with 123",
			args: args{int64(123)},
			want: nil,
			want1: &utils.ApplicationError{
				Message:    fmt.Sprintf("User with id = %v not found", int64(123)),
				StatusCode: http.StatusNotFound,
				Code:       "not found",
			},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			got, got1 := GetUser(tt.args.userID)
			if !reflect.DeepEqual(got, tt.want) {
				b.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				b.Errorf("GetUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
