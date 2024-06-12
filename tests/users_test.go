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
			name: "get list of users",
			in:   API,
			filter: &pkg.UserFilter{
				CustomerId: "test",
			},
			want: []pkg.User{},
		},
	}

	for _, test := range tests {
		users, _ := API.GetUsers(test.filter)
		if len(users) != len(test.want) {
			t.Errorf("%s - got %d lengh, want %d length", test.name, len(users), len(test.want))
		}
	}
}
