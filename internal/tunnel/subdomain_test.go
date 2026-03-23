package tunnel

import "testing"

func TestValidateSubdomain(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
	}{
		{"", false},
		{"myapp", false},
		{"cool-project", false},
		{"demo", false},
		{"paypal", true},
		{"pay-pal", true},
		{"paypal-login", true},
		{"my-google-app", true},
		{"apple", true},
		{"facebook", true},
		{"PAYPAL", true},
		{"Amazon", true},
		{"chase-bank", true},
		{"my-stripe-test", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			err := ValidateSubdomain(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSubdomain(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
