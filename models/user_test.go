package models

import "testing"

func TestUser_IsValid(t *testing.T) {
	tests := []struct {
		name string
		u    *User
		want bool
	}{
		{"nil", nil, false},
		{"empty struct", &User{}, false},
		{"no password", &User{Email: "test@test.nl"}, false},
		{"minimal requirments", &User{Email: "test@test.nl", Password: "geheim1234"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.IsValid(); got != tt.want {
				t.Errorf("User.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
