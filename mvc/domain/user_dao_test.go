package domain

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang-microservices/mvc/utils"
)

var (
	UserDaoMock     userDaoMock
	getUserFunction func(userID int64) (*User, *utils.ApplicationError)
)

type userDaoMock struct{}

func (m *userDaoMock) GetUser(userID int64) (*User, *utils.ApplicationError) {
	return m.GetUserFunction(userID)
}

func init() {
	UserDao = &userDaoMock{}
}

func (m *userDaoMock) GetUserFunction(userID int64) (*User, *utils.ApplicationError) {
	fmt.Println("inside mock function")
	switch userID {
	// user found test
	case 1:
		return &User{
			ID:        1,
			FirstName: "Eduardo",
			LastName:  "Fernandez",
			Email:     "fenandez9000@gmail.com",
		}, nil
	// user not found test
	case 123:
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("User with id = %v not found", userID),
			StatusCode: http.StatusNotFound,
			Code:       "not found",
		}
	}
	return nil, nil
}

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
			got, got1 := UserDao.GetUser(tt.args.userID)
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
			got, got1 := UserDao.GetUser(tt.args.userID)
			if !reflect.DeepEqual(got, tt.want) {
				b.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				b.Errorf("GetUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
