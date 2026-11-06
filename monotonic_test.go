package monotonic

import "testing"

func TestCheckUpgradeable(t *testing.T) {
	tests := []struct {
		name  string
		state EntityStabilityStateValue
		want  bool
	}{
		{
			name:  "current value is greater",
			state: EntityStabilityStateValue{Previous: 0.5, Current: 0.7},
			want:  true,
		},
		{
			name:  "current value is equal",
			state: EntityStabilityStateValue{Previous: 0.5, Current: 0.5},
			want:  true,
		},
		{
			name:  "current value is lower",
			state: EntityStabilityStateValue{Previous: 0.7, Current: 0.5},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckUpgradeable(tt.state); got != tt.want {
				t.Errorf("CheckUpgradeable(%+v) = %t, want %t", tt.state, got, tt.want)
			}
		})
	}
}
