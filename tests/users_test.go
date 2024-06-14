package tests

import (
	"testing"

	"github.com/pstano1/go-cart/internal/api"
	"github.com/pstano1/go-cart/internal/pkg"
)

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name   string
		in     *api.InstanceAPI
		filter *pkg.UserFilter
		want   []pkg.User
	}{
		{
			name: "non existent customerId",
			in:   API,
			filter: &pkg.UserFilter{
				CustomerId: "test",
			},
			want: make([]pkg.User, 0),
		},
		{
			name: "query all users for one customer",
			in:   API,
			filter: &pkg.UserFilter{
				CustomerId: customerId,
			},
			want: make([]pkg.User, 1),
		},
		{
			name: "query user by username (user doesn't exist)",
			in:   API,
			filter: &pkg.UserFilter{
				Username:   "testuser",
				CustomerId: customerId,
			},
			want: make([]pkg.User, 0),
		},
		{
			name: "query user by username (user exists)",
			in:   API,
			filter: &pkg.UserFilter{
				Username:   "admin",
				CustomerId: customerId,
			},
			want: make([]pkg.User, 1),
		},
		{
			name: "query user by id (user doesn't exist)",
			in:   API,
			filter: &pkg.UserFilter{
				Id:         "test",
				CustomerId: customerId,
			},
			want: make([]pkg.User, 0),
		},
		{
			name: "query user by id (user exists)",
			in:   API,
			filter: &pkg.UserFilter{
				Id:         "243277f5-146d-4dc1-9c66-11b41ced6ead",
				CustomerId: customerId,
			},
			want: make([]pkg.User, 1),
		},
	}

	for _, test := range tests {
		users, _ := test.in.GetUsers(test.filter)
		if len(users) != len(test.want) {
			t.Errorf("%s - got %d length, want %d length", test.name, len(users), len(test.want))
		}
	}
}

func TestGetPermissions(t *testing.T) {
	tests := []struct {
		name string
		in   *api.InstanceAPI
		want []string
	}{
		{
			name: "query permissions",
			in:   API,
			want: make([]string, 19),
		},
	}

	for _, test := range tests {
		permissions, _ := test.in.GetPermissions()
		if len(permissions) != len(test.want) {
			t.Errorf("%s - got %d length, want %d length", test.name, len(permissions), len(test.want))
		}
	}
}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name    string
		in      *api.InstanceAPI
		payload *pkg.UserCreate
		want    error
	}{
		{
			name: "proper user setup with taken username",
			in:   API,
			payload: &pkg.UserCreate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				Username:      "admin",
				Password:      "StrongPassword123",
				PasswordCheck: "StrongPassword123",
				Email:         "admin@example.com",
				Permissions:   make([]string, 0),
				IsActive:      true,
			},
			want: pkg.ErrUserAlreadyExists,
		},
		{
			name: "user setup with not matching passwords",
			in:   API,
			payload: &pkg.UserCreate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				Username:      "admin",
				Password:      "StrongPassword123",
				PasswordCheck: "StrongPassword12",
				Email:         "admin@example.com",
				Permissions:   make([]string, 0),
				IsActive:      true,
			},
			want: pkg.ErrPassowordsDontMatch,
		},
		{
			name: "proper user setup",
			in:   API,
			payload: &pkg.UserCreate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				Username:      "testuser",
				Password:      "StrongPassword123",
				PasswordCheck: "StrongPassword123",
				Email:         "admin@example.com",
				Permissions:   make([]string, 0),
				IsActive:      true,
			},
			want: nil,
		},
	}

	for _, test := range tests {
		_, err := test.in.CreateUser(test.payload)
		if err != test.want {
			t.Errorf("%s - got %d, want %d", test.name, err, test.want)
		}
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name    string
		in      *api.InstanceAPI
		payload *pkg.UserUpdate
		want    error
	}{
		{
			name: "non existent id in payload",
			in:   API,
			payload: &pkg.UserUpdate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				Email:       "admin@example.com",
				Permissions: make([]string, 0),
				Id:          "test",
			},
			want: pkg.ErrUserNotFound,
		},
		{
			name: "proper payload",
			in:   API,
			payload: &pkg.UserUpdate{
				CustomerSpecificModel: pkg.CustomerSpecificModel{
					CustomerId: customerId,
				},
				Email:       "admin@example.com",
				Permissions: make([]string, 0),
				Id:          "243277f5-146d-4dc1-9c66-11b41ced6ead",
			},
			want: nil,
		},
	}

	for _, test := range tests {
		err := test.in.UpdateUser(test.payload)
		if err != test.want {
			t.Errorf("%s - got %d, want %d", test.name, err, test.want)
		}
	}
}
