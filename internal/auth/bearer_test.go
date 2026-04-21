package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerFromHeader(t *testing.T) {
	tests := []struct{
		name 		string
		header		string
		token		string
		gotError	bool	
	}{
		{
			name:      "valid bearer token",
			header:    "Bearer peD+oiGKQf5ATtHoANHq9PwIKLCTvO3i+sQfKo1EKIIk1VXX62imu7ZXhlWWNVgZ",
			token: 	"peD+oiGKQf5ATtHoANHq9PwIKLCTvO3i+sQfKo1EKIIk1VXX62imu7ZXhlWWNVgZ",
			gotError:   false,
		},
		{
			name:    "empty header",
			header:  "",
			gotError: true,
		},
		{
			name:    "missing token",
			header:  "Bearer",
			gotError: true,
		},
		{
			name:    "wrong scheme",
			header:  "Briar peD+oiGKQf5ATtHoANHq9PwIKLCTvO3i+sQfKo1EKIIk1VXX62imu7ZXhlWWNVgZ",
			gotError: true,
		},
		{
			name:    "token without bearer prefix",
			header:  "peD+oiGKQf5ATtHoANHq9PwIKLCTvO3i+sQfKo1EKIIk1VXX62imu7ZXhlWWNVgZ",
			gotError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			header := http.Header{}

			if test.header != "" {
				header.Add("Authorization", test.header)
			}

			tokenFromFunction, err := GetBearerToken(header)

			if test.gotError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error occured %v", err)
			}

			if tokenFromFunction != test.token {
				t.Fatalf("expected output: %v, got %v", test.token, tokenFromFunction)
			}
		})
	}

}